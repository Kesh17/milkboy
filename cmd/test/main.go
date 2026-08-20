package main

import (
	"log/slog"
	"milkboy/bus"
	"milkboy/cartidge"
	"milkboy/cpu"
	"milkboy/gameboy"
	"time"
)

// only for texting purpose
func main() {
	// c, err := cartidge.New("roms/01-special.gb")
	slog.SetLogLoggerLevel(slog.LevelDebug)
	c, err := cartidge.New("roms/gb.gb")
	if err != nil {
		slog.Error("cartidge load error", "error: ", err)
	}
	bus := bus.New(c)
	cpu := cpu.New()

	gb := gameboy.New(cpu, bus)

	for {
		gb.Cycle()
		time.Sleep(time.Millisecond)
	}

}
