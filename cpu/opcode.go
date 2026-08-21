package cpu

import (
	"fmt"
	"log/slog"
)

type Instruction func(opcode uint8)

func setLDr8r8(c *CPU) {
	for i := 0x40; i <= 0x7F; i++ {
		if i == 0x76 {
			//todo: halt
			continue
		}
		c.OpcodeTable[i] = c.LDr8r8
	}
}

func (c *CPU) populateTable() {
	c.OpcodeTable[0x0] = c.NOP

	//ld r8 r8
	setLDr8r8(c)

}

func (c *CPU) NOP(opcode uint8) {
	slog.Debug("Decode NOP")
}

func (c *CPU) LDr8r8(opcode uint8) {
	src := R8(opcode & 0b00000111)
	dest := R8((opcode & 0b00111000) >> 3)
	slog.Debug("Decode LD r8,r8",
		"opcode", fmt.Sprintf("0x%02X", opcode),
		"src", src,
		"dest", dest,
	)
	c.SetR8(src, c.GetR8(dest))
}

func (c *CPU) LDr8n8(opcode uint8) {
	r8 := R8(opcode & 0b00111000)
	n8 := c.Fetch()

	c.SetR8(r8, n8)
}
