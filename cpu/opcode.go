package cpu

type Instruction func()

func (c *CPU) populateTable() {
	c.OpcodeTable[0x0] = c.NOP
}

func (c *CPU) NOP() {

}
