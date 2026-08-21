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
		//todo
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
