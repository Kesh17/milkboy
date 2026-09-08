package main

import (
	"log/slog"
	"milkboy/bus"
	"milkboy/cartridge"
	"milkboy/cpu"
	"milkboy/gameboy"
	"milkboy/interrupt"
	"milkboy/timer"
)

// only for texting purpose
func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	c, err := cartridge.New("roms/tetris.gb")
	if err != nil {
		slog.Error("cartridge load error", "error: ", err)
	}

	interrupt := &interrupt.Interrupt{}
	timer := timer.New(interrupt)
	bus := bus.New(c, interrupt, timer)
	cpu := cpu.New(bus, interrupt)

	gb := gameboy.New(cpu, bus, timer)

	gb.Run()
}
