package cpu

import "log/slog"

// all instructions returns the number of cycles they are taking
type Instruction func(opcode uint8) uint64

func setLDR8R8(c *CPU) {
	for i := 0x40; i <= 0x7F; i++ {
		if i == 0x76 {
			//todo: halt
			continue
		}
		c.opcodeTable[i] = c.ldR8R8
	}
}

func setLDR8N8(c *CPU) {
	for i := 0x06; i <= 0x3E; i += 0x08 {
		c.opcodeTable[i] = c.ldR8N8
	}
}

func setLDR16N16(c *CPU) {
	for i := 0x01; i <= 0x31; i += 0x10 {
		c.opcodeTable[i] = c.ldR16N16
	}
}

func setLDR16memA(c *CPU) {
	for i := 0x02; i <= 0x32; i += 0x10 {
		c.opcodeTable[i] = c.ldR16memA
	}
}

func setLDAR16mem(c *CPU) {
	for i := 0x0A; i <= 0x3A; i += 0x10 {
		c.opcodeTable[i] = c.ldAR16mem
	}
}

func setANDAR8(c *CPU) {
	for i := 0xA0; i <= 0xA7; i++ {
		c.opcodeTable[i] = c.andAR8
	}
}

func setORAR8(c *CPU) {
	for i := 0xB0; i <= 0xB7; i++ {
		c.opcodeTable[i] = c.orAR8
	}
}

func setXORAR8(c *CPU) {
	for i := 0xA8; i <= 0xAF; i++ {
		c.opcodeTable[i] = c.xorAR8
	}
}

func setBITU8R(c *CPU) {
	for i := 0x40; i <= 0x7F; i++ {
		c.opcodeTable[i] = c.bitU3R8
	}
}

func setRESU8R(c *CPU) {
	for i := 0x80; i <= 0xBF; i++ {
		c.opcodeTable[i] = c.resU3R8
	}
}

func setSETAR8(c *CPU) {
	for i := 0xC0; i <= 0xFF; i++ {
		c.opcodeTable[i] = c.setU3R8
	}
}

func (c *CPU) populateTable() {
	//nop
	c.opcodeTable[0x0] = c.NOP

	//ld r8 r8
	setLDR8R8(c)

	//ld r8 n8
	setLDR8N8(c)

	//ld r16 n16
	setLDR16N16(c)

	//ld r16mem a
	setLDR16memA(c)

	//ld [n16] a
	c.opcodeTable[0xEA] = c.ldN16A

	//ldh [n16] a
	c.opcodeTable[0xE0] = c.ldhN16A

	//ldh [c] a
	c.opcodeTable[0xE2] = c.ldhCA

	//ld a r16mem
	setLDAR16mem(c)

	//ld a [n16]
	c.opcodeTable[0xFA] = c.ldAN16

	// ldh a [n16]
	c.opcodeTable[0xF0] = c.ldhAN16

	//ldh a c
	c.opcodeTable[0xF2] = c.ldhAC

	//ldh [n16] sp
	c.opcodeTable[0x08] = c.ldN16SP

	//ld hl sp+e8
	c.opcodeTable[0xF8] = c.ldHLSPe8

	//ld sp hl
	c.opcodeTable[0xF9] = c.ldSPHL

	//AND A,r8
	setANDAR8(c)

	//AND A,n8
	c.opcodeTable[0xE6] = c.andAN8

	//OR A,r8
	setORAR8(c)

	// OR A,n8
	c.opcodeTable[0xF6] = c.orAN8

	//XOR A,r8
	setXORAR8(c)

	// XOR A,n8
	c.opcodeTable[0xF6] = c.xorAN8

	//BIT A, r8
	setBITU8R(c)

	//RES A, r8
	setRESU8R(c)

	//SET A, r8
	setSETAR8(c)

}

func (c *CPU) NOP(opcode uint8) uint64 {
	cycles := uint64(1)
	return cycles
}

func (c *CPU) ldR8R8(opcode uint8) uint64 {
	src := R8(opcode & 0b00000111)
	dest := R8((opcode & 0b00111000) >> 3)
	c.SetR8(dest, c.GetR8(src))

	var cycles uint64
	if src == R8HL || dest == R8HL {
		cycles = 2
	} else {
		cycles = 1
	}
	return cycles
}

func (c *CPU) ldR8N8(opcode uint8) uint64 {
	r8 := R8((opcode & 0b00111000) >> 3)
	n8 := c.Fetch()

	c.SetR8(r8, n8)

	var cycles uint64
	if r8 == R8HL {
		cycles = 3
	} else {
		cycles = 2
	}
	return cycles
}

func (c *CPU) ldR16N16(opcode uint8) uint64 {
	lo := c.Fetch()
	hi := c.Fetch()

	r16 := R16((opcode & 0b00110000) >> 4)
	n16 := uint16(hi)<<8 | uint16(lo)

	c.SetR16(r16, n16)

	cycles := uint64(3)
	return cycles
}

func (c *CPU) ldR16memA(opcode uint8) uint64 {
	r16 := R16Mem((opcode & 0b00110000) >> 4)
	c.SetR16Mem(r16, c.register.A)

	cycles := uint64(2)
	return cycles
}

func (c *CPU) ldN16A(opcode uint8) uint64 {
	lo := c.Fetch()
	hi := c.Fetch()
	n16 := uint16(hi)<<8 | uint16(lo)

	c.bus.Write(n16, c.register.A)

	cycles := uint64(4)
	return cycles
}

func (c *CPU) ldhN16A(opcode uint8) uint64 {
	lo := uint16(c.Fetch())
	hi := uint16(0xFF00)
	n16 := hi | lo

	c.bus.Write(n16, c.register.A)

	cycles := uint64(3)
	return cycles
}

func (c *CPU) ldhCA(opcode uint8) uint64 {
	addr := uint16(0xFF00) | uint16(c.register.C)

	c.bus.Write(addr, c.register.A)

	cycles := uint64(2)
	return cycles
}

func (c *CPU) ldAR16mem(opcode uint8) uint64 {
	r16 := R16Mem((opcode & 0b00110000) >> 4)
	data := c.GetR16Mem(r16)

	c.register.A = data

	cycles := uint64(2)
	return cycles
}

func (c *CPU) ldAN16(opcode uint8) uint64 {
	lo := c.Fetch()
	hi := c.Fetch()
	n16 := uint16(hi)<<8 | uint16(lo)

	c.register.A = c.bus.Read(n16)

	cycles := uint64(4)
	return cycles
}

func (c *CPU) ldhAN16(opcode uint8) uint64 {
	lo := c.Fetch()
	n16 := 0xFF00 | uint16(lo)

	c.register.A = c.bus.Read(n16)

	cycles := uint64(3)
	return cycles
}

func (c *CPU) ldhAC(opcode uint8) uint64 {
	addr := uint16(0xFF00) | uint16(c.register.C)

	c.register.A = c.bus.Read(addr)

	cycles := uint64(2)
	return cycles
}

func (c *CPU) ldN16SP(opcode uint8) uint64 {
	lo := c.Fetch()
	hi := c.Fetch()
	n16 := uint16(hi)<<8 | uint16(lo)

	low := c.sp & 0xFF
	high := c.sp >> 8
	c.bus.Write(n16, byte(low))
	c.bus.Write(n16+1, byte(high))

	cycles := uint64(5)
	return cycles
}

func (c *CPU) ldHLSPe8(opcode uint8) uint64 {
	e8 := int8(c.Fetch())
	sp := c.sp

	result := uint16(int32(sp) + int32(e8))

	halfCarry := ((sp & 0x0F) + (uint16(uint8(e8)) & 0x0F)) > 0x0F
	carry := ((sp & 0xFF) + uint16(uint8(e8))) > 0xFF

	c.register.H = uint8(result >> 8)
	c.register.L = uint8(result)

	c.register.F = 0
	if halfCarry {
		c.register.F.SetFlag(H)
	}
	if carry {
		c.register.F.SetFlag(C)
	}

	cycles := uint64(3)
	return cycles
}

func (c *CPU) ldSPHL(opcode uint8) uint64 {
	c.sp = c.register.HL()

	cycles := uint64(2)
	return cycles
}

func (c *CPU) andAR8(opcode uint8) uint64 {

	r8 := R8(opcode & 0x7)
	c.register.A &= c.GetR8(r8)

	if c.register.A == 0 {
		c.register.F.SetFlag(Z)
	} else {
		c.register.F.ClearFlag(Z)
	}

	c.register.F.ClearFlag(N)
	c.register.F.SetFlag(H)
	c.register.F.ClearFlag(C)

	var cycles uint64
	if r8 == R8HL {
		cycles = 2
	} else {
		cycles = 1
	}
	slog.Debug("Decode AND A,r8", "r8", r8)
	return cycles
}

func (c *CPU) andAN8(opcode uint8) uint64 {
	n8 := c.Fetch()
	c.register.A &= n8

	slog.Debug("Decode AND A,n8", "n8", n8)
	cycles := uint64(2)
	return cycles
}

func (c *CPU) cpl(opcode uint8) uint64 {
	c.register.A ^= c.register.A

	c.register.F.SetFlag(N)
	c.register.F.SetFlag(H)

	slog.Debug("Decode CPL")

	cycles := uint64(1)
	return cycles
}

func (c *CPU) orAR8(opcode uint8) uint64 {

	r8 := R8(opcode & 0x7)
	c.register.A |= c.GetR8(r8)

	if c.register.A == 0 {
		c.register.F.SetFlag(Z)
	} else {
		c.register.F.ClearFlag(Z)
	}

	c.register.F.ClearFlag(N)
	c.register.F.ClearFlag(H)
	c.register.F.ClearFlag(C)

	var cycles uint64
	if r8 == R8HL {
		cycles = 2
	} else {
		cycles = 1
	}
	slog.Debug("Decode OR A,r8", "r8", r8)
	return cycles
}

func (c *CPU) orAN8(opcode uint8) uint64 {
	n8 := c.Fetch()
	c.register.A |= n8

	slog.Debug("Decode OR A,n8", "n8", n8)
	cycles := uint64(2)
	return cycles
}

func (c *CPU) xorAR8(opcode uint8) uint64 {

	r8 := R8(opcode & 0x7)
	c.register.A ^= c.GetR8(r8)

	if c.register.A == 0 {
		c.register.F.SetFlag(Z)
	} else {
		c.register.F.ClearFlag(Z)
	}

	c.register.F.ClearFlag(N)
	c.register.F.ClearFlag(H)
	c.register.F.ClearFlag(C)

	var cycles uint64
	if r8 == R8HL {
		cycles = 2
	} else {
		cycles = 1
	}
	slog.Debug("Decode XOR A,r8", "r8", r8)
	return cycles
}

func (c *CPU) xorAN8(opcode uint8) uint64 {
	n8 := c.Fetch()
	c.register.A ^= n8

	slog.Debug("Decode XOR A,n8", "n8", n8)
	cycles := uint64(2)
	return cycles
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
