package cpu

import (
	"fmt"
	"log/slog"
	"milkboy/bus"
)

type CPU struct {
	Bus      bus.Addresable
	Register Register

	PC uint16
	SP uint16

	OpcodeTable [256]Instruction
}

func New(bus bus.Addresable) *CPU {
	c := &CPU{PC: 0x100, Bus: bus}
	c.populateTable()
	return c
}

func (c *CPU) Fetch() uint8 {
	debug_pc := c.PC
	opcode := c.Bus.Read(c.PC)
	c.PC++
	slog.Debug("Fetch",
		"PC", fmt.Sprintf("%#X", debug_pc),
		"opcode", fmt.Sprintf("%#X", opcode),
	)
	return opcode
}
