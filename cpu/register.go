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
