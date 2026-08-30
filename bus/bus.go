package bus

import (
	"fmt"
	"log/slog"
	"milkboy/cartridge"
	"milkboy/interrupt"
	"milkboy/io_register"
)

type Addresable interface {
	Read(uint16) byte
	Write(uint16, byte)
}

type Bus struct {
	cartridge *cartridge.Cartridge
	vram      [8 * 1024]byte
	io        io_register.IORegister
	wram      [8 * 1024]byte //todo: switchable banks for cgb mode
	hram      [127]byte
	interrupt *interrupt.Interrupt
}

func New(cart *cartridge.Cartridge) *Bus {
	//only for now
	i := &interrupt.Interrupt{}
	return &Bus{cartridge: cart, interrupt: i}
}

func (b *Bus) Write(addr uint16, data byte) {
	switch {
	case addr <= 0x7FFF:
		b.cartridge.Write(addr, data)

	case addr <= 0x9FFF:
		b.vram[addr-0x8000] = data

	case addr <= 0xBFFF:
		b.cartridge.WriteRam(addr, data)

	case addr <= 0xDFFF:
		b.wram[addr-0xC000] = data

	case addr <= 0xFDFF: //Echo ram

	case addr <= 0xFE9F: //OAM

	case addr <= 0xFEFF: //Not usable

	case addr <= 0xFF7F:
		b.io.Write(addr, data)

	case addr <= 0xFFFE: //HRAM not implemented dummy
		b.hram[addr-0xFF80] = data

	case addr == 0xFFFF:
		b.interrupt.Write(data)

	default:
		slog.Warn("couldn't find location to write to", "addr", fmt.Sprintf("0x%04X", addr))
	}

}

func (b *Bus) Read(addr uint16) byte {
	var data uint8
	switch {
	case addr <= 0x7FFF:
		return b.cartridge.Read(addr)

	case addr <= 0x9FFF:
		return b.vram[addr-0x8000]

	case addr <= 0xBFFF:
		return b.cartridge.Read(addr)

	case addr <= 0xDFFF:
		return b.wram[addr-0xC000]

	case addr <= 0xFDFF: //Echo ram

	case addr <= 0xFE9F: //OAM

	case addr <= 0xFEFF: //Not usable

	case addr <= 0xFF7F:
		return b.io.Read(addr)

	case addr <= 0xFFFE: //HRAM not implemented dummy
		return b.hram[addr-0xFF80]

	case addr == 0xFFFF:
		return b.interrupt.Read()

	default:
		slog.Warn("opcode will have garbage value")
	}
	return data
}
