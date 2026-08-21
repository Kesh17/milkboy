package io_register

import "log/slog"

type IORegister struct {
	serial Serial
}

type Serial struct {
	sb uint8
	sc uint8
}

func (ir *IORegister) Read(addr uint16) byte {
	slog.Debug("Reading from IO Register")
	var data uint8
	switch {
	//Serial Registers
	case addr >= 0xFF01 && addr <= 0xFF02:
		switch addr {
		case 0xFF01:
			data = ir.serial.sb
		case 0xFF02:
			data = ir.serial.sc
		}
	default:
		slog.Warn("this IO Register is not defined so garbage read")
	}

	return data
}
