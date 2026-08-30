package main

import (
	"log/slog"
	"milkboy/bus"
	"milkboy/cartidge"
	"milkboy/cpu"
	"milkboy/gameboy"
)

// only for texting purpose
func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	c, err := cartidge.New("roms/tetris.gb")
	if err != nil {
		slog.Error("cartidge load error", "error: ", err)
	}
	bus := bus.New(c)
	cpu := cpu.New(bus)

	gb := gameboy.New(cpu, bus)

	for {
		gb.CPU.Cycle()
		// time.Sleep(time.Millisecond)
	}

}
