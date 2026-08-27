package cpu

import "log/slog"

func populateCBopcodeTable(c *CPU) {
	//BIT A, r8
	setBITU8R(c)

	//RES A, r8
	setRESU8R(c)

	//SET A, r8
	setSETAR8(c)

	//RL r8
	setRLR8(c)

	//RLC r8
	setRLCR8(c)

	//RR r8
	setRRR8(c)

	//RRC r8
	setRRCR8(c)

	//SLA r8
	setSLAR8(c)

	//SRA r8
	setSRAR8(c)

	//SRL r8
	setSRLR8(c)

	//SWAP r8
	setSWAPR8(c)
}

func setRLR8(c *CPU) {
	for i := 0x10; i <= 0x17; i++ {
		c.cbInstructionTable[i] = c.rlR8
	}
}

func setRLCR8(c *CPU) {
	for i := 0x00; i <= 0x07; i++ {
		c.cbInstructionTable[i] = c.rlcR8
	}
}

func setRRR8(c *CPU) {
	for i := 0x18; i <= 0x1F; i++ {
		c.cbInstructionTable[i] = c.rrR8
	}
}

func setRRCR8(c *CPU) {
	for i := 0x08; i <= 0x0F; i++ {
		c.cbInstructionTable[i] = c.rrcR8
	}
}

func setSLAR8(c *CPU) {
	for i := 0x20; i <= 0x27; i++ {
		c.cbInstructionTable[i] = c.slaR8
	}
}

func setSRAR8(c *CPU) {
	for i := 0x28; i <= 0x2F; i++ {
		c.cbInstructionTable[i] = c.sraR8
	}
}

func setSRLR8(c *CPU) {
	for i := 0x38; i <= 0x3F; i++ {
		c.cbInstructionTable[i] = c.srlR8
	}
}

func setSWAPR8(c *CPU) {
	for i := 0x30; i <= 0x37; i++ {
		c.cbInstructionTable[i] = c.swapR8
	}
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

func (c *CPU) rlR8(opcode uint8) uint64 {
	r8 := R8(opcode & 0x07)
	value := c.GetR8(r8)

	newCarry := value&0x80 != 0

	oldCarry := uint8(0)
	if c.register.F.C() {
		oldCarry = 1
	}

	result := (value << 1) | oldCarry
	c.SetR8(r8, result)

	cycles := uint64(2)

	c.register.F.setFlagIf(Z, result == 0)
	c.register.F.ClearFlag(N)
	c.register.F.ClearFlag(H)
	c.register.F.setFlagIf(C, newCarry)

	if r8 == R8HL {
		cycles = 4
	}

	slog.Debug("Decode RL r8", "r8", r8)
	return cycles
}

func (c *CPU) rlcR8(opcode uint8) uint64 {
	r8 := R8(opcode & 0x07)
	value := c.GetR8(r8)
	newCarry := value&0x80 != 0
	value = (value << 1) | (value >> 7)
	c.SetR8(r8, value)

	c.register.F.setFlagIf(Z, value == 0)
	c.register.F.ClearFlag(N)
	c.register.F.ClearFlag(H)
	c.register.F.setFlagIf(C, newCarry)

	cycles := uint64(2)

	if r8 == R8HL {
		cycles = 4
	}

	slog.Debug("Decode RLC r8", "r8", r8)
	return cycles
}

func (c *CPU) rrR8(opcode uint8) uint64 {
	r8 := R8(opcode & 0x07)
	value := c.GetR8(r8)

	newCarry := value&0x1 != 0

	oldCarry := uint8(0)
	if c.register.F.C() {
		oldCarry = 1
	}

	result := (oldCarry << 7) | (value >> 1)
	c.SetR8(r8, result)

	cycles := uint64(2)

	c.register.F.setFlagIf(Z, result == 0)
	c.register.F.ClearFlag(N)
	c.register.F.ClearFlag(H)
	c.register.F.setFlagIf(C, newCarry)

	if r8 == R8HL {
		cycles = 4
	}

	slog.Debug("Decode RR r8", "r8", r8)
	return cycles
}

func (c *CPU) rrcR8(opcode uint8) uint64 {
	r8 := R8(opcode & 0x07)
	value := c.GetR8(r8)
	newCarry := value&0x01 != 0
	value = (value >> 1) | (value << 7)

	c.SetR8(r8, value)

	c.register.F.setFlagIf(Z, value == 0)
	c.register.F.ClearFlag(N)
	c.register.F.ClearFlag(H)
	c.register.F.setFlagIf(C, newCarry)

	cycles := uint64(2)

	if r8 == R8HL {
		cycles = 4
	}

	slog.Debug("Decode RRC r8", "r8", r8)
	return cycles
}

func (c *CPU) slaR8(opcode uint8) uint64 {
	r8 := R8(opcode & 0x07)
	value := c.GetR8(r8)
	newCarry := value&0x80 != 0
	value = (value << 1)

	c.SetR8(r8, value)

	c.register.F.setFlagIf(Z, value == 0)
	c.register.F.ClearFlag(N)
	c.register.F.ClearFlag(H)
	c.register.F.setFlagIf(C, newCarry)

	cycles := uint64(2)

	if r8 == R8HL {
		cycles = 4
	}

	slog.Debug("Decode SLA r8", "r8", r8)
	return cycles
}

func (c *CPU) sraR8(opcode uint8) uint64 {
	r8 := R8(opcode & 0x07)
	value := c.GetR8(r8)

	newCarry := value&0x01 != 0
	value = (value >> 1) | (value & 0x80)

	c.SetR8(r8, value)

	c.register.F.setFlagIf(Z, value == 0)
	c.register.F.ClearFlag(N)
	c.register.F.ClearFlag(H)
	c.register.F.setFlagIf(C, newCarry)

	cycles := uint64(2)

	if r8 == R8HL {
		cycles = 4
	}

	slog.Debug("Decode SRA r8", "r8", r8)
	return cycles
}

func (c *CPU) srlR8(opcode uint8) uint64 {
	r8 := R8(opcode & 0x07)
	value := c.GetR8(r8)
	newCarry := value&0x01 != 0
	value = value >> 1

	c.SetR8(r8, value)

	c.register.F.setFlagIf(Z, value == 0)
	c.register.F.ClearFlag(N)
	c.register.F.ClearFlag(H)
	c.register.F.setFlagIf(C, newCarry)

	cycles := uint64(2)

	if r8 == R8HL {
		cycles = 4
	}

	slog.Debug("Decode SRL r8", "r8", r8)
	return cycles
}

func (c *CPU) swapR8(opcode uint8) uint64 {
	r8 := R8(opcode & 0x07)
	value := c.GetR8(r8)
	value = (value << 4) | (value >> 4)
	c.SetR8(r8, value)

	c.register.F.setFlagIf(Z, value == 0)
	c.register.F.ClearFlag(N)
	c.register.F.ClearFlag(H)
	c.register.F.ClearFlag(C)

	cycles := uint64(2)

	if r8 == R8HL {
		cycles = 4
	}

	slog.Debug("Decode SWAP r8", "r8", r8)
	return cycles
}
