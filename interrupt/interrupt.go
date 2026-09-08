package interrupt

import (
	"log/slog"
)

type InterruptType byte

const (
	VBlank InterruptType = iota
	LCD
	Timer
	Serial
	Joypad
)

type Interrupt struct {
	ieBit byte
	ifBit byte
	IME   bool
}

func (i *Interrupt) Read(addr uint16) byte {
	if addr == 0xFF0F {
		return i.ifBit
	}
	return i.ieBit
}

func (i *Interrupt) Write(addr uint16, value byte) {
	if addr == 0xFF0F {
		i.ifBit = value
		return
	}
	i.ieBit = value
}

func (i *Interrupt) ClearIFBit(interruptType InterruptType) {
	i.ifBit &^= 1 << interruptType
}

func (i *Interrupt) SetIFBit(interruptType InterruptType) {
	i.ifBit |= 1 << interruptType
}

func (i *Interrupt) GetPending() (InterruptType, bool) {
	for j := range InterruptType(5) {
		if i.ifBit&(1<<j) != 0 && i.ieBit&(1<<j) != 0 {
			return j, true
		}
	}
	return 0, false
}

func (i *Interrupt) GetAddress(interruptType InterruptType) uint16 {
	switch interruptType {
	case VBlank:
		return 0x0040
	case LCD:
		return 0x0048
	case Timer:
		return 0x0050
	case Serial:
		return 0x0058
	case Joypad:
		return 0x0060
	default:
		slog.Warn("Unknown interrupt type")
		return 0
	}
}

func (i *Interrupt) Request(interruptType InterruptType) {
	i.SetIFBit(interruptType)
}

func (i *Interrupt) IsPending() bool {
	return i.ifBit&i.ieBit != 0
}
