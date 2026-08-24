package cpu

import (
	"fmt"
	"log/slog"
	"milkboy/bus"
)

type CPU struct {
	bus      bus.Addresable
	register Register

	pc uint16
	sp uint16

	ime bool

	instructionTable   [256]Instruction
	cbInstructionTable [256]Instruction
}

func New(bus bus.Addresable) *CPU {
	c := &CPU{pc: 0x100, bus: bus, sp: 0xFFFE}
	c.populateTable()
	return c
}

func (c *CPU) Cycle() {
	opcode := c.Fetch()
	c.DecodeExecute(opcode)
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
