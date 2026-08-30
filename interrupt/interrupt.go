package interrupt

type Interrupt struct {
	ie byte
}

func (i *Interrupt) Read() byte {
	return i.ie
}

func (i *Interrupt) Write(value byte) {
	i.ie = value
}
