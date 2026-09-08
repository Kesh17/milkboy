package gameboy

import (
	"milkboy/bus"
	"milkboy/cpu"
	"milkboy/timer"
)

type GameBoy struct {
	cpu   *cpu.CPU
	bus   *bus.Bus
	timer *timer.Timer
}

func New(cpu *cpu.CPU, bus *bus.Bus, timer *timer.Timer) *GameBoy {
	return &GameBoy{cpu: cpu, bus: bus, timer: timer}
}

func (gb *GameBoy) Run() {
	for {
		cycles := gb.cpu.Cycle() * 4
		for range cycles {
			gb.timer.Cycle()
		}
	}
}
