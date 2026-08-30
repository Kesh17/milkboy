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
			data = ir.serial.sb
		case 0xFF02:
			data = ir.serial.sc
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
			ir.serial.sb = data
		case 0xFF02:
			//for blargg's test
			if data == 0x81 {
				fmt.Printf("%c", ir.serial.sb)
			}
			ir.serial.sc = data
		}
	default:
		slog.Warn("this IO Register is not defined so nothing to write", "addr", addr)
	}
}
