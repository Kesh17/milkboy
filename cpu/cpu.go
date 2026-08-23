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

	opcodeTable [256]Instruction
}

func New(bus bus.Addresable) *CPU {
	c := &CPU{pc: 0x100, bus: bus}
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

func (c *CPU) DecodeExecute(opcode byte) {
	if c.opcodeTable[opcode] == nil {
		slog.Warn("Not implemented yet", "opcode", fmt.Sprintf("%02X", opcode))
		return
	}
	c.opcodeTable[opcode](opcode)

}
