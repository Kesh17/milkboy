package gameboy

import (
	"fmt"
	"log/slog"
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

func (gb *GameBoy) Fetch() byte {
	opcode := gb.Bus.Read(gb.CPU.PC)
	return opcode
}

func (gb *GameBoy) DecodeExecute(opcode byte) {
	if gb.CPU.OpcodeTable[opcode] == nil {
		slog.Warn("Not implemented yet", "opcode", fmt.Sprintf("%02X", opcode))
		return
	}
	gb.CPU.OpcodeTable[opcode](opcode)

}

func (gb *GameBoy) Cycle() {
	opcode := gb.CPU.Fetch()
	gb.DecodeExecute(opcode)
}
