package cpu

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
		return cycles
	} else {
		cycles = 1
		return cycles
	}
}

func (c *CPU) ldR8N8(opcode uint8) uint64 {
	r8 := R8((opcode & 0b00111000) >> 3)
	n8 := c.Fetch()

	c.SetR8(r8, n8)

	var cycles uint64
	if r8 == R8HL {
		cycles = 3
		return cycles
	} else {
		cycles = 2
		return cycles
	}
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
