package cpu

import "log/slog"

type FlagRegister uint8

type FlagBit uint8

const (
	Z FlagBit = 0x80
	N FlagBit = 0x40
	H FlagBit = 0x20
	C FlagBit = 0x10
)

const (
	FBNZ FlagBit = iota
	FBZ
	FBNC
	FBC
)

func (f FlagBit) String() string {
	switch f {
	case FBNZ:
		return "FBNZ"
	case FBZ:
		return "FBZ"
	case FBNC:
		return "FBNC"
	case FBC:
		return "FBC"
	default:
		return "unknown"
	}
}

func (f *FlagRegister) ConditionCode(fb FlagBit) bool {
	res := false

	switch fb {
	case FBNZ:
		return !f.isSet(Z)
	case FBZ:
		return f.isSet(Z)
	case FBNC:
		return !f.isSet(C)
	case FBC:
		return f.isSet(C)

	default:
		slog.Warn("unknown flag bit")
	}

	return res
}

func (f *FlagRegister) isSet(fb FlagBit) bool {
	return *f&FlagRegister(fb) != 0
}
func (f *FlagRegister) Z() bool {
	return f.isSet(Z)
}

func (f *FlagRegister) N() bool {
	return f.isSet(N)
}

func (f *FlagRegister) H() bool {
	return f.isSet(H)
}

func (f *FlagRegister) C() bool {
	return f.isSet(C)
}

func (f *FlagRegister) SetFlag(fb FlagBit) {
	*f |= FlagRegister(fb)
}

func (f *FlagRegister) ClearFlag(fb FlagBit) {
	*f &^= FlagRegister(fb)
}

func (f *FlagRegister) setFlagIf(flag FlagBit, cond bool) {
	if cond {
		f.SetFlag(flag)
	} else {
		f.ClearFlag(flag)
	}
}
