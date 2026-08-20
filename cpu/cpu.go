package cpu

type CPU struct {
	Register Register

	PC uint16
	SP uint16

	OpcodeTable [256]Instruction
}

func New() *CPU {
	c := &CPU{PC: 0x100}
	c.populateTable()
	return c
}
