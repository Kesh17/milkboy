package gameboy

import (
	"milkboy/bus"
	"milkboy/cpu"
)

type GameBoy struct {
	CPU *cpu.CPU
	Bus *bus.Bus
}

func New(cpu *cpu.CPU, bus *bus.Bus) *GameBoy {
	return &GameBoy{CPU: cpu, Bus: bus}
}
