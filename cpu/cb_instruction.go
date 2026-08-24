package cpu

import "log/slog"

func populateCBopcodeTable(c *CPU) {
	//BIT A, r8
	setBITU8R(c)

	//RES A, r8
	setRESU8R(c)

	//SET A, r8
	setSETAR8(c)
}

func setBITU8R(c *CPU) {
	for i := 0x40; i <= 0x7F; i++ {
		c.cbInstructionTable[i] = c.bitU3R8
	}
}

func setRESU8R(c *CPU) {
	for i := 0x80; i <= 0xBF; i++ {
		c.cbInstructionTable[i] = c.resU3R8
	}
}

func setSETAR8(c *CPU) {
	for i := 0xC0; i <= 0xFF; i++ {
		c.cbInstructionTable[i] = c.setU3R8
	}
}

func (c *CPU) bitU3R8(opcode uint8) uint64 {
	b3 := (opcode >> 3) & 0x7
	r8 := R8(opcode & 0x7)
	value := c.GetR8(r8)

	if value&(1<<b3) == 0 {
		c.register.F.SetFlag(Z)
	} else {
		c.register.F.ClearFlag(Z)
	}
	c.register.F.ClearFlag(N)
	c.register.F.SetFlag(H)

	slog.Debug("Decode BIT U3,r8", "b3", b3, "r8", r8)

	var cycles uint64
	if r8 == R8HL {
		cycles = 3
	} else {
		cycles = 2
	}
	return cycles
}

func (c *CPU) resU3R8(opcode uint8) uint64 {
	b3 := (opcode >> 3) & 0x7
	mask := uint8(1) << b3
	r8 := R8(opcode & 0x7)
	value := c.GetR8(r8)
	c.SetR8(r8, value&^mask)

	slog.Debug("Decode RES U3,r8", "b3", b3, "r8", r8)

	var cycles uint64
	if r8 == R8HL {
		cycles = 4
	} else {
		cycles = 2
	}
	return cycles
}
func (c *CPU) setU3R8(opcode uint8) uint64 {
	b3 := (opcode >> 3) & 0x7
	mask := uint8(1) << b3
	r8 := R8(opcode & 0x7)
	value := c.GetR8(r8)
	c.SetR8(r8, value|mask)

	slog.Debug("Decode SET U3,r8", "b3", b3, "r8", r8)

	var cycles uint64
	if r8 == R8HL {
		cycles = 4
	} else {
		cycles = 2
	}
	return cycles
}
