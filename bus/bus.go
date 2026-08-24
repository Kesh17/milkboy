package bus

import (
	"fmt"
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
	wram     [0x2000]byte
}

func New(cart *cartidge.Cartidge) *Bus {
	return &Bus{Cartidge: cart}
}

func (b *Bus) Write(addr uint16, data byte) {
	switch {
	case addr <= 0x7FFF:
		b.Cartidge.Write(addr, data)
	case addr >= 0xC000 && addr <= 0xDFFF:
		b.wram[addr-0xC000] = data
	case addr >= 0xE000 && addr <= 0xFDFF:
		b.wram[addr-0xE000] = data
	case addr >= 0xFF00 && addr <= 0xFF7F:
		b.io.Write(addr, data)
	default:
		slog.Warn("couldn't find location to write to", "addr", fmt.Sprintf("0x%04X", addr))
	}

}

func (b *Bus) Read(addr uint16) byte {
	var data uint8
	switch {
	case addr <= 0x7FFF:
		data = b.Cartidge.Read(addr)
	case addr >= 0xC000 && addr <= 0xDFFF:
		return b.wram[addr-0xC000]
	case addr >= 0xFF00 && addr <= 0xFF7F:
		data = b.io.Read(addr)
	default:
		slog.Warn("opcode will have garbage value")
	}
	return data
}
