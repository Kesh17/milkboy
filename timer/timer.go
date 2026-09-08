package timer

import (
	"log/slog"
	"milkboy/interrupt"
)

type Timer struct {
	systemCounter uint16

	tima byte
	tma  byte
	tac  byte

	interrupt *interrupt.Interrupt
}

func New(interrupt *interrupt.Interrupt) *Timer {
	return &Timer{interrupt: interrupt}
}

func (t *Timer) Cycle() {
	old := t.timerSignal()

	t.systemCounter++

	new := t.timerSignal()

	if old && !new {
		t.incrementTima()
	}
}

func (t *Timer) timerSignal() bool {
	if t.tac&0x04 == 0 {
		return false
	}

	var bit uint16
	switch t.tac & 0x03 {
	case 0x0:
		bit = 9
	case 0x1:
		bit = 3
	case 0x2:
		bit = 5
	case 0x3:
		bit = 7
	default:
		slog.Warn("Invalid timer frequency", "frequency", t.tac&0x03)
		return false
	}

	return (t.systemCounter & (1 << bit)) != 0
}

func (t *Timer) incrementTima() {
	if t.tima == 0xFF {
		t.tima = t.tma
		t.interrupt.SetIFBit(interrupt.Timer)
		return
	}

	t.tima++
}

func (t *Timer) ReadDiv() byte {
	return byte(t.systemCounter >> 8)
}

func (t *Timer) Read(addr uint16) byte {
	switch addr {
	case 0xFF04:
		return t.ReadDiv()
	case 0xFF05:
		return t.tima
	case 0xFF06:
		return t.tma
	case 0xFF07:
		return t.tac
	default:
		slog.Warn("Invalid timer read address", "address", addr)
		return 0
	}
}

func (t *Timer) Write(addr uint16, value byte) {
	switch addr {
	case 0xFF04:
		t.resetDIV()
	case 0xFF05:
		t.tima = value
	case 0xFF06:
		t.tma = value
	case 0xFF07:
		t.writeTAC(value)
	default:
		slog.Warn("Invalid timer write address", "address", addr)
	}
}

func (t *Timer) writeTAC(value byte) {
	old := t.timerSignal()

	t.tac = value & 0x07

	new := t.timerSignal()

	if old && !new {
		t.incrementTima()
	}
}

func (t *Timer) resetDIV() {
	old := t.timerSignal()

	t.systemCounter = 0

	new := t.timerSignal()

	if old && !new {
		t.incrementTima()
	}
}
