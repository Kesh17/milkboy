package bus

import (
	"log/slog"
	"milkboy/cartidge"
	io "milkboy/io_register"
)

type Addresable interface {
	Read(uint16) byte
	Write(uint16, byte)
}

type Bus struct {
	Cartidge *cartidge.Cartidge
	io       io.IORegister
}

func New(cart *cartidge.Cartidge) *Bus {
	return &Bus{Cartidge: cart}
}

func (b *Bus) Write(addr uint16, data byte) {
	switch {
	case addr <= 0x7FFF:
		b.Cartidge.Write(addr, data)

	case addr >= 0xFF00 && addr <= 0xFF7F:
		b.io.Write(addr, data)
	default:
		slog.Warn("couldn't find location to write to")
	}

}

func (b *Bus) Read(addr uint16) byte {
	var data uint8
	switch {
	case addr <= 0x7FFF:
		data = b.Cartidge.Read(addr)

	case addr >= 0xFF00 && addr <= 0xFF7F:
		data = b.io.Read(addr)
	default:
		slog.Warn("opcode will have garbage value")
	}

	return data
}
