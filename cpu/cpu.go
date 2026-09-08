package cpu

import (
	"fmt"
	"log/slog"
	"milkboy/bus"
	"milkboy/interrupt"
)

type CPU struct {
	bus      bus.Addresable
	register Register

	pc uint16
	sp uint16

	interrupt *interrupt.Interrupt
	pending   bool // for EI instruction, to enable interrupts after next instruction
	halted    bool
	haltBug   bool

	instructionTable   [256]Instruction
	cbInstructionTable [256]Instruction
}

func New(bus bus.Addresable, interrupt *interrupt.Interrupt) *CPU {
	f := FlagRegister(0)
	f.SetFlag(Z)
	f.ClearFlag(N)
	//reading cartidge header checksum
	if bus.Read(0x14D) == 0x00 {
		f.ClearFlag(H)
		f.ClearFlag(C)
	} else {
		f.SetFlag(H)
		f.SetFlag(C)
	}

	c := &CPU{
		register: Register{
			A: 0x01, F: f, B: 0x00, C: 0x13, D: 0x00, E: 0xD8, H: 0x01, L: 0x4D,
		},
		pc: 0x100, sp: 0xFFFE, bus: bus, interrupt: interrupt}

	c.populateTable()
	c.populateCBopcodeTable()
	return c
}

func (c *CPU) Cycle() uint64 {
	cycles := c.handleInterrupts()
	if c.halted {
		return 1
	}

	//halt bug not implemented
	opcode := c.Fetch()
	cycles += c.DecodeExecute(opcode)

	if c.pending {
		c.pending = false
		c.interrupt.IME = true
	}

	return cycles
}

func (c *CPU) Fetch() uint8 {
	debug_pc := c.pc
	opcode := c.bus.Read(c.pc)
	c.pc++
	slog.Debug("Fetch",
		"PC", fmt.Sprintf("%#X", debug_pc),
		"opcode", fmt.Sprintf("%#X", opcode),
	)
	return opcode
}

func (c *CPU) Fetch16() uint16 {
	lo := c.Fetch()
	hi := c.Fetch()
	n16 := uint16(hi)<<8 | uint16(lo)

	return n16
}

func (c *CPU) DecodeExecute(opcode byte) uint64 {
	if opcode == 0xCB {
		opcode = c.Fetch()
		if c.cbInstructionTable[opcode] == nil {
			slog.Warn("Not implemented yet cb instruction", "opcode", fmt.Sprintf("%02X", opcode))
			return 0
		}
		return c.cbInstructionTable[opcode](opcode)
	}
	if c.instructionTable[opcode] == nil {
		slog.Warn("Not implemented yet", "opcode", fmt.Sprintf("%02X", opcode))
		return 0
	}
	return c.instructionTable[opcode](opcode)
}

func (c *CPU) push(value uint16) {
	c.sp--
	c.bus.Write(c.sp, uint8(value>>8))

	c.sp--
	c.bus.Write(c.sp, uint8(value))
}

func (c *CPU) pop() uint16 {
	lo := c.bus.Read(c.sp)
	c.sp++
	hi := c.bus.Read(c.sp)
	c.sp++
	return uint16(hi)<<8 | uint16(lo)
}

func (c *CPU) handleInterrupts() uint64 {
	pendingInterrupt, ok := c.interrupt.GetPending()
	if !ok {
		return 0
	}

	if c.halted {
		c.halted = false
	}

	if !c.interrupt.IME {
		return 0
	}

	c.serviceInterrupt(pendingInterrupt)
	return 5
}

func (c *CPU) serviceInterrupt(interruptType interrupt.InterruptType) {
	slog.Debug("Servicing interrupt", "type", interruptType)
	c.interrupt.ClearIFBit(interruptType)
	c.interrupt.IME = false
	c.push(c.pc)
	c.pc = c.interrupt.GetAddress(interruptType)
}
