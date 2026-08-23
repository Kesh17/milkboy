package cpu

type Register struct {
	A uint8
	F FlagRegister

	B uint8
	C uint8

	D uint8
	E uint8

	H uint8
	L uint8
}

type R8 uint8

const (
	R8B R8 = iota
	R8C
	R8D
	R8E
	R8H
	R8L
	R8HL
	R8A
)

func (r R8) String() string {
	switch r {
	case R8B:
		return "B"
	case R8C:
		return "C"
	case R8D:
		return "D"
	case R8E:
		return "E"
	case R8H:
		return "H"
	case R8L:
		return "L"
	case R8HL:
		return "[HL]"
	case R8A:
		return "A"
	default:
		return "UNKNOWN"
	}
}

func (c *CPU) SetR8(r8 R8, data uint8) {
	switch r8 {
	case R8B:
		c.register.B = data
	case R8C:
		c.register.C = data
	case R8D:
		c.register.D = data
	case R8E:
		c.register.E = data
	case R8H:
		c.register.H = data
	case R8L:
		c.register.L = data
	case R8HL:
		c.bus.Write(c.register.HL(), data)
	case R8A:
		c.register.A = data
	}
}

func (c *CPU) GetR8(r8 R8) uint8 {
	var data uint8
	switch r8 {
	case R8B:
		data = c.register.B
	case R8C:
		data = c.register.C
	case R8D:
		data = c.register.D
	case R8E:
		data = c.register.E
	case R8H:
		data = c.register.H
	case R8L:
		data = c.register.L
	case R8HL:
		data = c.bus.Read(c.register.HL())
	case R8A:
		data = c.register.A
	}

	return data
}

type R16 uint8

const (
	R16BC R16 = iota
	R16DE
	R16HL
	R16SP
)

func (r R16) String() string {
	switch r {
	case R16BC:
		return "BC"
	case R16DE:
		return "DE"
	case R16HL:
		return "HL"
	case R16SP:
		return "SP"
	default:
		return "UNKNOWN"
	}
}

func (c *CPU) SetR16(r16 R16, data uint16) {
	switch r16 {
	case R16BC:
		c.register.B = uint8(data >> 8)
		c.register.C = uint8(data)
	case R16DE:
		c.register.D = uint8(data >> 8)
		c.register.E = uint8(data)
	case R16HL:
		c.register.H = uint8(data >> 8)
		c.register.L = uint8(data)
	case R16SP:
		c.sp = data
	}
}

func (c *CPU) GetR16(r16 R16) uint16 {
	var data uint16
	switch r16 {
	case R16BC:
		data = c.register.BC()
	case R16DE:
		data = c.register.DE()
	case R16HL:
		data = c.register.HL()
	case R16SP:
		data = c.sp
	}
	return data
}

type R16Mem uint8

const (
	R16MemBC R16Mem = iota
	R16MemDE
	R16MemHLI
	R16MemHLD
)

func (r R16Mem) String() string {
	switch r {
	case R16MemBC:
		return "BC"
	case R16MemDE:
		return "DE"
	case R16MemHLI:
		return "HL+"
	case R16MemHLD:
		return "HL-"
	default:
		return "UNKNOWN"
	}
}

func (c *CPU) SetR16Mem(r16m R16Mem, data uint8) {
	switch r16m {
	case R16MemBC:
		c.bus.Write(c.register.BC(), data)
	case R16MemDE:
		c.bus.Write(c.register.DE(), data)
	case R16MemHLI:
		c.bus.Write(c.register.HL(), data)
		hl := c.register.HL()
		hl++
		c.register.H = uint8(hl >> 8)
		c.register.L = uint8(hl)
	case R16MemHLD:
		c.bus.Write(c.register.HL(), data)
		hl := c.register.HL()
		hl--
		c.register.H = uint8(hl >> 8)
		c.register.L = uint8(hl)
	}
}

func (c *CPU) GetR16Mem(r16 R16Mem) byte {
	var data byte
	switch r16 {
	case R16MemBC:
		data = c.bus.Read(c.register.BC())
	case R16MemDE:
		data = c.bus.Read(c.register.DE())
	case R16MemHLI:
		data = c.bus.Read(c.register.HL())
		hl := c.register.HL()
		hl++
		c.register.H = uint8(hl >> 8)
		c.register.H = uint8(hl)
	case R16MemHLD:
		data = c.bus.Read(c.register.HL())
		hl := c.register.HL()
		hl--
		c.register.H = uint8(hl >> 8)
		c.register.H = uint8(hl)
	}
	return data
}

func combineRegister(hi uint8, lo uint8) uint16 {
	return uint16(hi)<<8 | uint16(lo)
}

func (r *Register) AF() uint16 {
	return combineRegister(r.A, uint8(r.F))
}

func (r *Register) BC() uint16 {
	return combineRegister(r.B, r.C)
}

func (r *Register) DE() uint16 {
	return combineRegister(r.D, r.E)
}

func (r *Register) HL() uint16 {
	return combineRegister(r.H, r.L)
}
