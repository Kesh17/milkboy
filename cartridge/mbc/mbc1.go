package mbc

import (
	"log/slog"
)

type MBC1 struct {
	romBankMask uint8

	rom []byte
	ram []byte

	ramEnabled bool

	lowerBank   uint8
	higherBank  uint8
	bankingMode bool
}

func NewMBC1(rom []byte, noOfRombanks uint16, ramSize byte) *MBC1 {
	mbc := &MBC1{
		romBankMask: uint8(noOfRombanks - 1),
		rom:         rom,
		ram:         make([]byte, getRamSize(ramSize)),
		lowerBank:   1,
	}

	slog.Info("MBC1 initialized", "rom banks: ", noOfRombanks, "ram size: ", len(mbc.ram))

	return mbc
}

func (m *MBC1) Read(addr uint16) byte {
	switch {
	case addr <= 0x3FFF:
		bank := m.getLowerRomBank()
		offset := uint32(bank)*0x4000 + uint32(addr)
		return m.rom[offset]

	case addr <= 0x7FFF:
		romBank := m.getRomBank()
		offset := uint32(romBank)*0x4000 + uint32(addr-0x4000)
		return m.rom[offset]

	case addr >= 0xA000 && addr <= 0xBFFF:
		return m.ReadRam(addr)
	}

	return 0xFF
}

func (m *MBC1) Write(addr uint16, data byte) {
	switch {
	case addr <= 0x1FFF:
		m.ramEnabled = (data&0x0F == 0x0A)

	case addr <= 0x3FFF:
		m.lowerBank = (data & 0x1F)
		if m.lowerBank == 0 {
			m.lowerBank = 1
		}
	case addr <= 0x5FFF:
		m.higherBank = (data & 0x03)
	case addr <= 0x7FFF:
		m.bankingMode = (data & 0x01) == 1
	}
}

func (m *MBC1) WriteRam(addr uint16, data byte) {
	if !m.ramEnabled || len(m.ram) == 0 {
		return
	}
	bank := m.getRamBank()
	offset := uint32(bank)*0x2000 + uint32(addr-0xA000)
	m.ram[offset] = data
}

func (m *MBC1) ReadRam(addr uint16) byte {
	if !m.ramEnabled || len(m.ram) == 0 {
		return 0xFF
	}

	bank := m.getRamBank()
	offset := uint32(bank)*0x2000 + uint32(addr-0xA000)

	return m.ram[offset]
}

func (m *MBC1) getRomBank() uint8 {
	var bank uint8
	if !m.bankingMode {
		bank = (m.higherBank << 5) | m.lowerBank
	} else {
		bank = m.lowerBank
	}

	return bank & m.romBankMask
}

func (m *MBC1) getLowerRomBank() uint8 {
	if m.bankingMode {
		return m.higherBank << 5
	}
	return 0
}

func (m *MBC1) getRamBank() uint8 {
	if m.bankingMode {
		return m.higherBank
	}
	return 0
}

func getRamSize(sizeCode byte) int {
	switch sizeCode {
	case 0x00, 0x01:
		return 0
	case 0x02:
		return 0x2000
	case 0x03:
		return 0x8000
	case 0x04:
		return 0x20000
	case 0x05:
		return 0x10000
	default:
		return 0
	}
}
