package io_register

import (
	"fmt"
	"log/slog"
)

type IORegister struct {
	serial Serial
}

type Serial struct {
	sb uint8
	sc uint8
}

func (ir *IORegister) Read(addr uint16) byte {
	var data uint8
	switch {
	//Serial Registers
	case addr >= 0xFF01 && addr <= 0xFF02:
		switch addr {
		case 0xFF01:
			slog.Debug("Reading from serial Register sb")
			data = ir.serial.sb
		case 0xFF02:
			slog.Debug("Reading from serial Register sc")
			data = ir.serial.sc
			// if data&0x80 != 0 {
			// serial transfer requested
			fmt.Printf("Output: %c", ir.serial.sb)
			// }
		}
	default:
		slog.Warn("this IO Register is not defined so garbage read", "addr", addr)
	}

	return data
}

func (ir *IORegister) Write(addr uint16, data byte) {
	switch {
	//Serial Registers
	case addr >= 0xFF01 && addr <= 0xFF02:
		switch addr {
		case 0xFF01:
			slog.Debug("Writing to serial Register sb")
			ir.serial.sb = data
		case 0xFF02:
			slog.Debug("Writing to serial Register sc")
			ir.serial.sc = data
		}
	default:
		slog.Warn("this IO Register is not defined so nothing to write to")
	}
}
