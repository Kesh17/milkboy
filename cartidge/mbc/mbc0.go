package mbc

type MBC0 struct {
	rom []byte
}

func NewMBC0(rom []byte) MBC {
	return &MBC0{rom: rom}
}

func (m *MBC0) Read(addr uint16) byte {
	return m.rom[addr]
}

func (m *MBC0) Write(addr uint16, data byte) {
	m.rom[addr] = data
}

// WriteRam is a no-op for MBC0 since it doesn't support external RAM
func (m *MBC0) WriteRam(addr uint16, data byte) {}
