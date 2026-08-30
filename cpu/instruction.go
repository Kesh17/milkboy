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
		c.instructionTable[i] = c.ldR8R8
	}
}

func setLDR8N8(c *CPU) {
	for i := 0x06; i <= 0x3E; i += 0x08 {
		c.instructionTable[i] = c.ldR8N8
	}
}

func setLDR16N16(c *CPU) {
	for i := 0x01; i <= 0x31; i += 0x10 {
		c.instructionTable[i] = c.ldR16N16
	}
}

func setLDR16memA(c *CPU) {
	for i := 0x02; i <= 0x32; i += 0x10 {
		c.instructionTable[i] = c.ldR16memA
	}
}

func setLDAR16mem(c *CPU) {
	for i := 0x0A; i <= 0x3A; i += 0x10 {
		c.instructionTable[i] = c.ldAR16mem
	}
}

func setANDAR8(c *CPU) {
	for i := 0xA0; i <= 0xA7; i++ {
		c.instructionTable[i] = c.andAR8
	}
}

func setORAR8(c *CPU) {
	for i := 0xB0; i <= 0xB7; i++ {
		c.instructionTable[i] = c.orAR8
	}
}

func setXORAR8(c *CPU) {
	for i := 0xA8; i <= 0xAF; i++ {
		c.instructionTable[i] = c.xorAR8
	}
}

func setCALLCCN16(c *CPU) {
	for i := 0xC4; i <= 0xDC; i += 0x08 {
		c.instructionTable[i] = c.callCCN16
	}
}

func setJPCCN16(c *CPU) {
	for i := 0xC2; i <= 0xDA; i += 0x08 {
		c.instructionTable[i] = c.jpCCN16
	}
}

func setJRCCN8(c *CPU) {
	for i := 0x20; i <= 0x38; i += 0x08 {
		c.instructionTable[i] = c.jrCCN8
	}
}

func setRST(c *CPU) {
	for i := 0xC7; i <= 0xFF; i += 0x08 {
		c.instructionTable[i] = c.rst
	}
}

func setRETCC(c *CPU) {
	for i := 0xC0; i <= 0xD8; i += 0x08 {
		c.instructionTable[i] = c.retCC
	}
}

func setADCAR8(c *CPU) {
	for i := 0x88; i <= 0x8F; i++ {
		c.instructionTable[i] = c.adcAR8
	}
}

func setADDAR8(c *CPU) {
	for i := 0x80; i <= 0x87; i++ {
		c.instructionTable[i] = c.addAR8
	}
}

func setCPAR8(c *CPU) {
	for i := 0xB8; i <= 0xBF; i++ {
		c.instructionTable[i] = c.cpAR8
	}
}

func setDECR8(c *CPU) {
	for i := 0x05; i <= 0x3D; i += 0x08 {
		c.instructionTable[i] = c.decR8
	}
}

func setINCAR8(c *CPU) {
	for i := 0x04; i <= 0x3c; i += 0x08 {
		c.instructionTable[i] = c.incR8
	}
}

func setSBCAR8(c *CPU) {
	for i := 0x98; i <= 0x9F; i++ {
		c.instructionTable[i] = c.sbcAR8
	}
}

func setSUBAR8(c *CPU) {
	for i := 0x90; i <= 0x97; i++ {
		c.instructionTable[i] = c.subAR8
	}
}

func setADDHLR16(c *CPU) {
	for i := 0x09; i <= 0x39; i += 0x10 {
		c.instructionTable[i] = c.addHLR16
	}
}

func setDECR16(c *CPU) {
	for i := 0x0B; i <= 0x3B; i += 0x10 {
		c.instructionTable[i] = c.decR16
	}
}

func setINCR16(c *CPU) {
	for i := 0x03; i <= 0x33; i += 0x10 {
		c.instructionTable[i] = c.incR16
	}
}

func setPUSHR16Stk(c *CPU) {
	for i := 0xC5; i <= 0xF5; i += 0x10 {
		c.instructionTable[i] = c.pushR16Stk
	}
}

func setPOPR16Stk(c *CPU) {
	for i := 0xC1; i <= 0xF1; i += 0x10 {
		c.instructionTable[i] = c.popR16Stk
	}
}

func (c *CPU) populateTable() {
	//nop
	c.instructionTable[0x0] = c.NOP

	//ld r8 r8
	setLDR8R8(c)

	//ld r8 n8
	setLDR8N8(c)

	//ld r16 n16
	setLDR16N16(c)

	//ld r16mem a
	setLDR16memA(c)

	//ld [n16] a
	c.instructionTable[0xEA] = c.ldN16A

	//ldh [n16] a
	c.instructionTable[0xE0] = c.ldhN16A

	//ldh [c] a
	c.instructionTable[0xE2] = c.ldhCA

	//ld a r16mem
	setLDAR16mem(c)

	//ld a [n16]
	c.instructionTable[0xFA] = c.ldAN16

	// ldh a [n16]
	c.instructionTable[0xF0] = c.ldhAN16

	//ldh a c
	c.instructionTable[0xF2] = c.ldhAC

	//ldh [n16] sp
	c.instructionTable[0x08] = c.ldN16SP

	//ld hl sp+e8
	c.instructionTable[0xF8] = c.ldHLSPe8

	//ld sp hl
	c.instructionTable[0xF9] = c.ldSPHL

	//AND A,r8
	setANDAR8(c)

	//AND A,n8
	c.instructionTable[0xE6] = c.andAN8

	//cpl
	c.instructionTable[0x2F] = c.cpl

	//OR A,r8
	setORAR8(c)

	// OR A,n8
	c.instructionTable[0xF6] = c.orAN8

	//XOR A,r8
	setXORAR8(c)

	// XOR A,n8
	c.instructionTable[0xEE] = c.xorAN8

	//CALL n16
	c.instructionTable[0xCD] = c.callN16

	//CALL CC n16
	setCALLCCN16(c)

	//JP HL
	c.instructionTable[0xE9] = c.jpHL

	//JP n16
	c.instructionTable[0xC3] = c.jpN16

	//JP CC n16
	setJPCCN16(c)

	//JR n8
	c.instructionTable[0x18] = c.jrN8

	//JR CC n8
	setJRCCN8(c)

	//RET CC
	setRETCC(c)

	//RET
	c.instructionTable[0xC9] = c.ret

	//RETI
	c.instructionTable[0xD9] = c.reti

	//RST
	setRST(c)

	//ADC A,r8
	setADCAR8(c)

	//ADC A,n8
	c.instructionTable[0xCE] = c.adcAN8

	//ADD A,r8
	setADDAR8(c)

	//ADD A,n8
	c.instructionTable[0xC6] = c.addAN8

	//CP A,r8
	setCPAR8(c)

	//CP A,n8
	c.instructionTable[0xFE] = c.cpAN8

	//DEC r8
	setDECR8(c)

	//INC r8
	setINCAR8(c)

	//SBC A R8
	setSBCAR8(c)

	//SBC A,n8
	c.instructionTable[0xDE] = c.sbcAN8

	//SUB A,r8
	setSUBAR8(c)

	//SUB A,n8
	c.instructionTable[0xD6] = c.subAN8

	//ADD HL, R16
	setADDHLR16(c)

	//DEC R16
	setDECR16(c)

	//INC R16
	setINCR16(c)

	//CCF
	c.instructionTable[0x3F] = c.ccf

	//SCF
	c.instructionTable[0x37] = c.scf

	//ADD SP e8
	c.instructionTable[0xE8] = c.addSPE8

	//PUSH r16stk
	setPUSHR16Stk(c)

	//POP r16stk
	setPOPR16Stk(c)

	//RLA
	c.instructionTable[0x17] = c.rla

	//RLCA
	c.instructionTable[0x07] = c.rlca

	//RRA
	c.instructionTable[0x1F] = c.rra

	//RRCA
	c.instructionTable[0x0F] = c.rrca

	//DAA
	c.instructionTable[0x27] = c.daa

	//DI
	c.instructionTable[0xF3] = c.di
}

func (c *CPU) NOP(opcode uint8) uint64 {
	cycles := uint64(1)

	slog.Debug("Decode NOP", "opcode", opcode)
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

	slog.Debug("Decode LD r8,r8", "r8", src, "r8", dest)

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

	slog.Debug("Decode LD r8,n8", "r8", r8, "n8", n8)
	return cycles
}

func (c *CPU) ldR16N16(opcode uint8) uint64 {
	lo := c.Fetch()
	hi := c.Fetch()

	r16 := R16((opcode & 0b00110000) >> 4)
	n16 := uint16(hi)<<8 | uint16(lo)

	c.SetR16(r16, n16)

	cycles := uint64(3)

	slog.Debug("Decode LD r16,n16", "r16", r16, "n16", n16)
	return cycles
}

func (c *CPU) ldR16memA(opcode uint8) uint64 {
	r16mem := R16Mem((opcode & 0b00110000) >> 4)
	c.SetR16Mem(r16mem, c.register.A)

	cycles := uint64(2)

	slog.Debug("Decode LD r16mem, A", "r16mem", r16mem)
	return cycles
}

func (c *CPU) ldN16A(opcode uint8) uint64 {
	n16 := c.Fetch16()

	c.bus.Write(n16, c.register.A)

	cycles := uint64(4)

	slog.Debug("Decode LD n16, A", "n16", n16)
	return cycles
}

func (c *CPU) ldhN16A(opcode uint8) uint64 {
	lo := c.Fetch()
	n16 := 0xFF00 | uint16(lo)

	c.bus.Write(n16, c.register.A)

	cycles := uint64(3)

	slog.Debug("Decode LD n16, A", "n16", n16)
	return cycles
}

func (c *CPU) ldhCA(opcode uint8) uint64 {
	addr := uint16(0xFF00) | uint16(c.register.C)

	c.bus.Write(addr, c.register.A)

	cycles := uint64(2)

	slog.Debug("Decode LD [C], A", "[c]", c.register.C)
	return cycles
}

func (c *CPU) ldAR16mem(opcode uint8) uint64 {
	r16mem := R16Mem((opcode & 0b00110000) >> 4)
	data := c.GetR16Mem(r16mem)

	c.register.A = data

	cycles := uint64(2)

	slog.Debug("Decode LD A, r16mem", "r16mem", r16mem)
	return cycles
}

func (c *CPU) ldAN16(opcode uint8) uint64 {
	n16 := c.Fetch16()

	c.register.A = c.bus.Read(n16)

	cycles := uint64(4)

	slog.Debug("Decode LD A,n16", "n16", n16)
	return cycles
}

func (c *CPU) ldhAN16(opcode uint8) uint64 {
	lo := c.Fetch()
	n16 := 0xFF00 | uint16(lo)

	c.register.A = c.bus.Read(n16)

	cycles := uint64(3)

	slog.Debug("Decode LDH A,n16", "n16", n16)
	return cycles
}

func (c *CPU) ldhAC(opcode uint8) uint64 {
	addr := uint16(0xFF00) | uint16(c.register.C)

	c.register.A = c.bus.Read(addr)

	cycles := uint64(2)

	slog.Debug("Decode LDH A, C", "addr", addr)
	return cycles
}

func (c *CPU) ldN16SP(opcode uint8) uint64 {
	n16 := c.Fetch16()

	low := c.sp & 0xFF
	high := c.sp >> 8
	c.bus.Write(n16, byte(low))
	c.bus.Write(n16+1, byte(high))

	cycles := uint64(5)

	slog.Debug("Decode LD n16, SP", "n16", n16)
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

	c.register.F.ClearFlag(Z)
	c.register.F.ClearFlag(N)
	c.register.F.setFlagIf(H, halfCarry)
	c.register.F.setFlagIf(C, carry)

	cycles := uint64(3)

	slog.Debug("Decode LD HL, SPe8", "spe8", result)
	return cycles
}

func (c *CPU) ldSPHL(opcode uint8) uint64 {
	c.sp = c.register.HL()

	cycles := uint64(2)

	slog.Debug("Decode LD SP, HL", "hl", c.register.HL())
	return cycles
}

func (c *CPU) andAR8(opcode uint8) uint64 {

	r8 := R8(opcode & 0x7)
	c.register.A &= c.GetR8(r8)

	c.register.F.setFlagIf(Z, c.register.A == 0)

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

	if c.register.A == 0 {
		c.register.F.SetFlag(Z)
	} else {
		c.register.F.ClearFlag(Z)
	}
	c.register.F.ClearFlag(N)
	c.register.F.SetFlag(H)
	c.register.F.ClearFlag(C)

	cycles := uint64(2)

	slog.Debug("Decode AND A,n8", "n8", n8)
	return cycles
}

func (c *CPU) cpl(opcode uint8) uint64 {
	c.register.A = ^c.register.A

	c.register.F.SetFlag(N)
	c.register.F.SetFlag(H)

	cycles := uint64(1)

	slog.Debug("Decode CPL")
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

	if c.register.A == 0 {
		c.register.F.SetFlag(Z)
	} else {
		c.register.F.ClearFlag(Z)
	}
	c.register.F.ClearFlag(N)
	c.register.F.ClearFlag(H)
	c.register.F.ClearFlag(C)

	cycles := uint64(2)

	slog.Debug("Decode OR A,n8", "n8", n8)
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

	if c.register.A == 0 {
		c.register.F.SetFlag(Z)
	} else {
		c.register.F.ClearFlag(Z)
	}

	c.register.F.ClearFlag(N)
	c.register.F.ClearFlag(H)
	c.register.F.ClearFlag(C)

	cycles := uint64(2)

	slog.Debug("Decode XOR A,n8", "n8", n8)
	return cycles
}
func (c *CPU) callN16(opcode uint8) uint64 {
	n16 := c.Fetch16()
	c.push(c.pc)
	c.pc = n16

	cycles := uint64(6)

	slog.Debug("Decode CALL n16", "n16", n16)
	return cycles
}

func (c *CPU) callCCN16(opcode uint8) uint64 {
	var cycles uint64
	cc := FlagBit((opcode >> 3) & 0x3)
	n16 := c.Fetch16()

	if !c.register.F.ConditionCode(cc) {
		cycles = 3
		return cycles
	}

	c.push(c.pc)
	c.pc = n16

	cycles = 6

	slog.Debug("Decode CALL cc n16", "cc", cc, "n16", n16)
	return cycles
}

func (c *CPU) jpHL(opcode uint8) uint64 {
	hl := c.register.HL()
	c.pc = hl

	cycles := uint64(1)

	slog.Debug("Decode JP HL", "hl", hl)
	return cycles
}

func (c *CPU) jpN16(opcode uint8) uint64 {
	n16 := c.Fetch16()
	c.pc = n16

	cycles := uint64(4)

	slog.Debug("Decode JP n16", "n16", n16)
	return cycles
}

func (c *CPU) jpCCN16(opcode uint8) uint64 {
	var cycles uint64
	cc := FlagBit((opcode >> 3) & 0x3)
	n16 := c.Fetch16()

	if !c.register.F.ConditionCode(cc) {
		cycles = 3
		return cycles
	}

	c.pc = n16

	cycles = 4

	slog.Debug("Decode JP cc n16", "cc", cc, "n16", n16)
	return cycles
}

func (c *CPU) jrN8(opcode uint8) uint64 {
	n8 := int8(c.Fetch())
	addr := int32(c.pc) + int32(n8)
	c.pc = uint16(addr)

	cycles := uint64(3)

	slog.Debug("Decode JR n8", "n8", n8)
	return cycles
}

func (c *CPU) jrCCN8(opcode uint8) uint64 {
	cc := FlagBit((opcode >> 3) & 0x3)
	n8 := int8(c.Fetch())
	var cycles uint64

	if !c.register.F.ConditionCode(cc) {
		cycles = 2
		return cycles
	}

	addr := uint16(int32(c.pc) + int32(n8))
	c.pc = addr

	cycles = 3

	slog.Debug("Decode JR cc n8", "cc", cc, "n8", n8)
	return cycles
}

func (c *CPU) retCC(opcode uint8) uint64 {
	var cycles uint64
	cc := FlagBit((opcode >> 3) & 0x3)

	if !c.register.F.ConditionCode(cc) {
		cycles = 2
		return cycles
	}

	c.pc = c.pop()

	cycles = 5

	slog.Debug("Decode RET cc", "cc", cc)
	return cycles
}

func (c *CPU) ret(opcode uint8) uint64 {
	c.pc = c.pop()

	cycles := uint64(4)

	slog.Debug("Decode RET")
	return cycles
}

func (c *CPU) reti(opcode uint8) uint64 {
	c.pc = c.pop()
	c.ime = true

	cycles := uint64(4)

	slog.Debug("Interrupt not implemented")
	slog.Debug("Decode RETI", "ime", c.ime)
	return cycles
}

func (c *CPU) rst(opcode uint8) uint64 {
	vec := ((opcode >> 3) & 0x7) * 8
	c.push(c.pc)
	c.pc = uint16(vec)

	cycles := uint64(4)

	slog.Debug("Decode RST", "vec", vec)
	return cycles
}

func (c *CPU) adcAR8(opcode uint8) uint64 {
	r8 := R8(opcode & 0x07)
	a := c.register.A
	value := c.GetR8(r8)

	var carryPlus uint8
	if c.register.F.C() {
		carryPlus = 1
	}

	res := uint16(a) + uint16(value) + uint16(carryPlus)
	c.register.A = uint8(res)

	halfCarry := (a&0xF)+(value&0xF)+carryPlus > 0xF
	carry := res > 0xFF

	c.register.F.setFlagIf(Z, res == 0)
	c.register.F.ClearFlag(N)
	c.register.F.setFlagIf(H, halfCarry)
	c.register.F.setFlagIf(C, carry)

	var cycles uint64
	if r8 == R8HL {
		cycles = 2
	} else {
		cycles = 1
	}

	slog.Debug("Decode ADC A r8", "r8", r8)
	return cycles
}

func (c *CPU) adcAN8(opcode uint8) uint64 {
	n8 := c.Fetch()
	a := c.register.A

	var carryPlus uint8
	if c.register.F.C() {
		carryPlus = 1
	}

	res := uint16(a) + uint16(n8) + uint16(carryPlus)
	c.register.A = uint8(res)

	halfCarry := (a&0x0F)+(n8&0x0F)+carryPlus > 0x0F
	carry := res > 0xFF

	c.register.F.setFlagIf(Z, res == 0)
	c.register.F.ClearFlag(N)
	c.register.F.setFlagIf(H, halfCarry)
	c.register.F.setFlagIf(C, carry)

	cycles := uint64(2)

	slog.Debug("Decode ADC A n8", "n8", n8)
	return cycles
}

func (c *CPU) addAR8(opcode uint8) uint64 {
	r8 := R8(opcode & 0x07)
	a := c.register.A
	value := c.GetR8(r8)
	result := a + value
	c.register.A = result

	c.register.F.setFlagIf(Z, result == 0)
	c.register.F.ClearFlag(N)
	c.register.F.setFlagIf(H, (a&0x0F)+(value&0x0F) > 0x0F)
	c.register.F.setFlagIf(C, uint16(a)+uint16(value) > 0xFF)

	cycles := uint64(1)
	if r8 == R8HL {
		cycles = 2
	}

	slog.Debug("Decode ADD A r8", "r8", r8)
	return cycles
}

func (c *CPU) addAN8(opcode uint8) uint64 {
	n8 := c.Fetch()
	a := c.register.A
	result := a + n8
	c.register.A = result

	c.register.F.setFlagIf(Z, result == 0)
	c.register.F.ClearFlag(N)
	c.register.F.setFlagIf(H, (a&0x0F)+(n8&0x0F) > 0x0F)
	c.register.F.setFlagIf(C, uint16(a)+uint16(n8) > 0xFF)

	cycles := uint64(1)

	slog.Debug("Decode ADD A n8", "r8", n8)
	return cycles
}

func (c *CPU) cpAR8(opcode uint8) uint64 {
	r8 := R8(opcode & 0x07)
	a := c.register.A
	value := c.GetR8(r8)
	result := a - value

	c.register.F.setFlagIf(Z, result == 0)
	c.register.F.SetFlag(N)
	c.register.F.setFlagIf(H, (a&0x0F) < (value&0x0F))
	c.register.F.setFlagIf(C, value > a)

	cycles := uint64(1)
	if r8 == R8HL {
		cycles = 2
	}

	slog.Debug("Decode CP A r8", "r8", r8)
	return cycles
}

func (c *CPU) cpAN8(opcode uint8) uint64 {
	n8 := c.Fetch()
	a := c.register.A
	result := a - n8

	c.register.F.setFlagIf(Z, result == 0)
	c.register.F.SetFlag(N)
	c.register.F.setFlagIf(H, (a&0x0F) < (n8&0x0F))
	c.register.F.setFlagIf(C, n8 > a)

	cycles := uint64(2)

	slog.Debug("Decode CP A n8", "n8", n8)
	return cycles
}

func (c *CPU) decR8(opcode uint8) uint64 {
	r8 := R8((opcode >> 3) & 0x07)
	value := c.GetR8(r8)
	result := value - 1
	c.SetR8(r8, result)

	c.register.F.setFlagIf(Z, result == 0)
	c.register.F.SetFlag(N)
	c.register.F.setFlagIf(H, value&0x0F == 0)

	cycles := uint64(1)
	if r8 == R8HL {
		cycles = 3
	}

	slog.Debug("Decode DEC r8", "r8", r8)
	return cycles
}

func (c *CPU) incR8(opcode uint8) uint64 {
	r8 := R8((opcode >> 3) & 0x07)
	result := c.GetR8(r8) + 1
	c.SetR8(r8, result)

	c.register.F.setFlagIf(Z, result == 0)
	c.register.F.ClearFlag(N)
	c.register.F.setFlagIf(H, result&0x0F == 0)

	cycles := uint64(1)
	if r8 == R8HL {
		cycles = 3
	}

	slog.Debug("DECODE INC r8", "r8", r8)
	return cycles
}

func (c *CPU) sbcAR8(opcode uint8) uint64 {
	r8 := R8(opcode & 0x07)
	a := c.register.A
	value := c.GetR8(r8)

	carryFlag := uint8(0)
	if c.register.F.C() {
		carryFlag = 1
	}

	result := a - value - carryFlag

	c.register.F.setFlagIf(Z, result == 0)
	c.register.F.SetFlag(N)
	c.register.F.setFlagIf(H, (a&0x0F) < ((value&0x0F)+carryFlag))
	c.register.F.setFlagIf(C, uint16(value)+uint16(carryFlag) > uint16(a))

	cycles := uint64(1)
	if r8 == R8HL {
		cycles = 2
	}

	slog.Debug("Decode SBC A r8", "r8", r8)
	return cycles
}

func (c *CPU) sbcAN8(opcode uint8) uint64 {
	n8 := c.Fetch()
	a := c.register.A

	carryFlag := uint8(0)
	if c.register.F.C() {
		carryFlag = 1
	}

	result := a - n8 - carryFlag

	c.register.F.setFlagIf(Z, result == 0)
	c.register.F.SetFlag(N)
	c.register.F.setFlagIf(H, (a&0x0F) < ((n8&0x0F)+carryFlag))
	c.register.F.setFlagIf(C, uint16(n8)+uint16(carryFlag) > uint16(a))

	cycles := uint64(2)

	slog.Debug("Decode SBC A n8", "n8", n8)
	return cycles
}

func (c *CPU) subAR8(opcode uint8) uint64 {
	r8 := R8(opcode & 0x07)
	a := c.register.A
	value := c.GetR8(r8)

	result := a - value
	c.register.A = result

	c.register.F.setFlagIf(Z, result == 0)
	c.register.F.SetFlag(N)
	c.register.F.setFlagIf(H, a&0x0F < value&0x0F)
	c.register.F.setFlagIf(C, value > a)

	cycles := uint64(1)
	if r8 == R8HL {
		cycles = 2
	}

	slog.Debug("Decode SUB A r8", "r8", r8)
	return cycles
}

func (c *CPU) subAN8(opcode uint8) uint64 {
	n8 := c.Fetch()
	a := c.register.A

	result := a - n8
	c.register.A = result

	c.register.F.setFlagIf(Z, result == 0)
	c.register.F.SetFlag(N)
	c.register.F.setFlagIf(H, a&0x0F < n8&0x0F)
	c.register.F.setFlagIf(C, n8 > a)

	cycles := uint64(2)

	slog.Debug("Decode SUB A n8", "n8", n8)
	return cycles
}

func (c *CPU) addHLR16(opcode uint8) uint64 {
	r16 := R16((opcode >> 4) & 0x03)
	hl := c.register.HL()
	value := c.GetR16(r16)
	result := hl + value
	c.SetR16(R16HL, result)

	c.register.F.ClearFlag(N)
	c.register.F.setFlagIf(H, (hl&0x0FFF)+(value&0x0FFF) > 0x0FFF)
	c.register.F.setFlagIf(C, uint32(hl)+uint32(value) > 0xFFFF)

	cycles := uint64(2)

	slog.Debug("Decode ADD HL r16", "r16", r16)
	return cycles
}

func (c *CPU) decR16(opcode uint8) uint64 {
	r16 := R16((opcode >> 4) & 0x03)
	value := c.GetR16(r16)
	c.SetR16(r16, value-1)

	cycles := uint64(2)

	slog.Debug("Decode DEC r16", "r16", r16)
	return cycles
}

func (c *CPU) incR16(opcode uint8) uint64 {
	r16 := R16((opcode >> 4) & 0x03)
	value := c.GetR16(r16)
	c.SetR16(r16, value+1)

	cycles := uint64(2)

	slog.Debug("Decode INC r16", "r16", r16)
	return cycles
}

func (c *CPU) ccf(opcode uint8) uint64 {
	c.register.F.ClearFlag(N)
	c.register.F.ClearFlag(H)
	c.register.F.setFlagIf(C, !c.register.F.C())

	cycles := uint64(1)
	slog.Debug("Decode CCF", "opcode", opcode)
	return cycles
}

func (c *CPU) scf(opcode uint8) uint64 {
	c.register.F.ClearFlag(N)
	c.register.F.ClearFlag(H)
	c.register.F.SetFlag(C)

	cycles := uint64(1)
	slog.Debug("Decode SCF", "opcode", opcode)
	return cycles
}

func (c *CPU) addSPE8(opcode uint8) uint64 {
	e8 := c.Fetch()
	sp := c.sp
	offset := int8(e8)
	result := uint16(int32(sp) + int32(offset))

	c.register.F.ClearFlag(Z)
	c.register.F.ClearFlag(N)
	c.register.F.setFlagIf(H, (sp&0x0F)+(uint16(e8)&0x0F) > 0x0F)
	c.register.F.setFlagIf(C, (sp&0xFF)+(uint16(e8)&0xFF) > 0xFF)

	c.sp = result

	cycles := uint64(1)
	slog.Debug("Decode ADD SP e8", "e8", e8)
	return cycles
}

func (c *CPU) popR16Stk(opcode uint8) uint64 {
	r16stk := R16Stk((opcode >> 4) & 0x03)

	sp := c.sp
	lo := c.bus.Read(sp)
	hi := c.bus.Read(sp + 1)
	c.sp = sp + 2

	value := (uint16(hi) << 8) | uint16(lo)
	c.setR16Stk(r16stk, value)

	cycles := uint64(3)
	slog.Debug("Decode POP r16stk", "r16stk", r16stk)
	return cycles
}

func (c *CPU) pushR16Stk(opcode uint8) uint64 {
	r16Stk := R16Stk((opcode >> 4) & 0x03)
	value := c.GetR16Stk(r16Stk)

	hi := byte(value >> 8)
	lo := byte(value)

	c.sp--
	c.bus.Write(c.sp, hi)

	c.sp--
	c.bus.Write(c.sp, lo)

	cycles := uint64(4)
	slog.Debug("Decode PUSH r16stk", "r16stk", r16Stk)
	return cycles
}

func (c *CPU) rla(opcode uint8) uint64 {
	value := c.register.A
	newCarry := value&0x80 != 0

	var oldCarry uint8
	if c.register.F.C() {
		oldCarry = 1
	}

	c.register.A = (value << 1) | oldCarry

	c.register.F.ClearFlag(Z)
	c.register.F.ClearFlag(N)
	c.register.F.ClearFlag(H)
	c.register.F.setFlagIf(C, newCarry)

	slog.Debug("Decode RLA", "A", c.register.A)
	cycles := uint64(1)
	return cycles
}

func (c *CPU) rlca(opcode uint8) uint64 {
	value := c.register.A
	newCarry := value&0x80 != 0

	c.register.A = (value << 1) | (value >> 7)

	c.register.F.ClearFlag(Z)
	c.register.F.ClearFlag(N)
	c.register.F.ClearFlag(H)
	c.register.F.setFlagIf(C, newCarry)

	slog.Debug("Decode RLCA", "A", c.register.A)
	cycles := uint64(1)
	return cycles
}

func (c *CPU) rra(opcode uint8) uint64 {
	value := c.register.A
	newCarry := value&0x01 != 0

	var oldCarry uint8
	if c.register.F.C() {
		oldCarry = 1
	}

	c.register.A = (oldCarry << 7) | (value >> 1)

	c.register.F.ClearFlag(Z)
	c.register.F.ClearFlag(N)
	c.register.F.ClearFlag(H)
	c.register.F.setFlagIf(C, newCarry)

	slog.Debug("Decode RRCA", "A", c.register.A)
	cycles := uint64(1)
	return cycles
}

func (c *CPU) rrca(opcode uint8) uint64 {
	value := c.register.A
	newCarry := value&0x01 != 0

	c.register.A = (value >> 1) | (value << 7)

	c.register.F.ClearFlag(Z)
	c.register.F.ClearFlag(N)
	c.register.F.ClearFlag(H)
	c.register.F.setFlagIf(C, newCarry)

	cycles := uint64(1)

	slog.Debug("Decode RRC A", "A", c.register.A)
	return cycles
}

func (c *CPU) di(opcode uint8) uint64 {
	c.ime = false

	cycles := uint64(1)

	slog.Debug("Interrupt not implemented")
	slog.Debug("Decode DI", "ime", c.ime)
	return cycles
}

func (c *CPU) daa(opcode uint8) uint64 {
	a := c.register.A
	carry := c.register.F.C()
	halfCarry := c.register.F.H()

	if !c.register.F.N() {
		adjustment := uint8(0)

		if halfCarry || a&0x0F > 0x09 {
			adjustment |= 0x06
		}
		if carry || a > 0x99 {
			adjustment |= 0x60
			carry = true
		}
		a += adjustment
	} else {
		adjustment := uint8(0)

		if halfCarry {
			adjustment |= 0x06
		}
		if carry {
			adjustment |= 0x60
		}
		a -= adjustment
	}

	c.register.A = a

	c.register.F.setFlagIf(Z, a == 0)
	c.register.F.ClearFlag(H)
	c.register.F.setFlagIf(C, carry)

	slog.Debug("Decode DAA", "A", c.register.A)
	cycles := uint64(1)
	return cycles
}
