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
		c.Register.B = data
	case R8C:
		c.Register.C = data
	case R8D:
		c.Register.D = data
	case R8E:
		c.Register.E = data
	case R8H:
		c.Register.H = data
	case R8L:
		c.Register.L = data
	case R8HL:
		c.Bus.Write(c.Register.HL(), data)
	case R8A:
		c.Register.A = data
	}
}

func (c *CPU) GetR8(r8 R8) uint8 {
	var data uint8
	switch r8 {
	case R8B:
		data = c.Register.B
	case R8C:
		data = c.Register.C
	case R8D:
		data = c.Register.D
	case R8E:
		data = c.Register.E
	case R8H:
		data = c.Register.H
	case R8L:
		data = c.Register.L
	case R8HL:
		data = c.Bus.Read(c.Register.HL())
	case R8A:
		data = c.Register.A
	}

	return data
}

type R16 uint16

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
		c.Register.B = uint8(data >> 8)
		c.Register.C = uint8(data)
	case R16DE:
		c.Register.D = uint8(data >> 8)
		c.Register.E = uint8(data)
	case R16HL:
		c.Register.H = uint8(data >> 8)
		c.Register.L = uint8(data)
	case R16SP:
		c.SP = data
	}
}

func (c *CPU) GetR16(r16 R16) uint16 {
	var data uint16
	switch r16 {
	case R16BC:
		data = c.Register.BC()
	case R16DE:
		data = c.Register.DE()
	case R16HL:
		data = c.Register.HL()
	case R16SP:
		data = c.SP
	}
	return data
}

type R16Mem uint16

const (
	R16M_BC R16Mem = iota
	R16M_DE
	R16M_HLI
	R16M_HLD
)

func (r R16Mem) String() string {
	switch r {
	case R16M_BC:
		return "BC"
	case R16M_DE:
		return "DE"
	case R16M_HLI:
		return "HL+"
	case R16M_HLD:
		return "HL-"
	default:
		return "UNKNOWN"
	}
}

func (c *CPU) SetR16Mem(r16m R16Mem, data uint8) {
	switch r16m {
	case R16M_BC:
		c.Bus.Write(uint16(r16m), data)
	case R16M_DE:
		c.Bus.Write(uint16(r16m), data)
	case R16M_HLI:
		c.Bus.Write(c.Register.HL(), data)
		hl := c.Register.HL()
		hl++
		c.Register.H = uint8(hl >> 8)
		c.Register.H = uint8(hl)
	case R16M_HLD:
		c.Bus.Write(c.Register.HL(), data)
		hl := c.Register.HL()
		hl--
		c.Register.H = uint8(hl >> 8)
		c.Register.H = uint8(hl)
	}
}

func (c *CPU) GetR16Mem(r16 R16Mem) byte {
	var data byte
	switch r16 {
	case R16M_BC:
		data = c.Bus.Read(c.Register.BC())
	case R16M_DE:
		data = c.Bus.Read(c.Register.DE())
	case R16M_HLI:
		data = c.Bus.Read(c.Register.HL())
		hl := c.Register.HL()
		hl--
		c.Register.H = uint8(hl >> 8)
		c.Register.H = uint8(hl)
	case R16M_HLD:
		data = c.Bus.Read(c.Register.HL())
		hl := c.Register.HL()
		hl++
		c.Register.H = uint8(hl >> 8)
		c.Register.H = uint8(hl)
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
