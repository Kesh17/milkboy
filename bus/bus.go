package bus

import (
	"log/slog"
	"milkboy/cartidge"
)

type Addresable interface {
	Read(uint8) byte
}

type Bus struct {
	Cartidge *cartidge.Cartidge
}

func New(cart *cartidge.Cartidge) *Bus {
	return &Bus{Cartidge: cart}
}

func (b *Bus) Read(addr uint16) byte {
	var opcode uint8
	switch {
	case addr <= 0x7FFF:
		opcode = b.Cartidge.Read(addr)
	default:
		slog.Warn("opcode will have garbage value")
	}

	return opcode
}
