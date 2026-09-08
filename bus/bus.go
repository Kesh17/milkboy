package bus

import (
	"fmt"
	"log/slog"
	"milkboy/cartridge"
	"milkboy/interrupt"
	"milkboy/io_register"
	"milkboy/timer"
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
	timer     *timer.Timer
}

func New(cart *cartridge.Cartridge, interrupt *interrupt.Interrupt, timer *timer.Timer) *Bus {
	return &Bus{cartridge: cart, interrupt: interrupt, timer: timer}
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
		switch addr {
		case 0xFF04, 0xFF05, 0xFF06, 0xFF07:
			b.timer.Write(addr, data)

		case 0xFF0F:
			b.interrupt.Write(addr, data)

		default:
			b.io.Write(addr, data)
		}

	case addr <= 0xFFFE: //HRAM not implemented dummy
		b.hram[addr-0xFF80] = data

	case addr == 0xFFFF:
		b.interrupt.Write(addr, data)

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
		switch addr {
		case 0xFF04, 0xFF05, 0xFF06, 0xFF07:
			return b.timer.Read(addr)

		case 0xFF0F:
			return b.interrupt.Read(addr)

		default:
			return b.io.Read(addr)
		}

	case addr <= 0xFFFE: //HRAM not implemented dummy
		return b.hram[addr-0xFF80]

	case addr == 0xFFFF:
		return b.interrupt.Read(addr)

	default:
		slog.Warn("opcode will have garbage value")
	}
	return data
}
