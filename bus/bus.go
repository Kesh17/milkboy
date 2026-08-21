package bus

import (
	"log/slog"
	"milkboy/cartidge"
	io "milkboy/io_register"
)

type Addresable interface {
	Read(uint16) byte
}

type Bus struct {
	Cartidge *cartidge.Cartidge
	io       io.IORegister
}

func New(cart *cartidge.Cartidge) *Bus {
	return &Bus{Cartidge: cart}
}

func (b *Bus) Read(addr uint16) byte {
	var opcode uint8
	switch {
	case addr <= 0x7FFF:
		opcode = b.Cartidge.Read(addr)

	case addr >= 0xFF00 && addr <= 0xFF7F:
		opcode = b.io.Read(addr)
	default:
		slog.Warn("opcode will have garbage value")
	}

	return opcode
}
