package cartridge

import (
	"log/slog"
	"milkboy/cartridge/mbc"
	"os"
	"strings"
)

type Cartridge struct {
	mbc.MBC
	Title    string
	filename string
}

func New(path string) (*Cartridge, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return &Cartridge{}, err
	}

	c := &Cartridge{filename: path}

	switch data[0x0147] {
	case 0x00:
		c.MBC = mbc.NewMBC0(data)
	default:
		slog.Warn("Unknown MBC type")
	}

	c.setHeaderChecksum()
	c.setTitle()
	return c, nil
}

func (c *Cartridge) headerChecksum() byte {
	var checksum byte = 0
	for address := uint16(0x0134); address <= 0x014C; address++ {
		checksum += c.Read(address)
	}
	return checksum
}

func (c *Cartridge) setHeaderChecksum() {
	c.Write(0x014D, c.headerChecksum())
}

func (c *Cartridge) setTitle() {
	var title strings.Builder
	for address := uint16(0x0134); address <= 0x0143; address++ {
		byte := c.Read(address)
		if byte == 0x00 {
			break
		}
		title.WriteByte(byte)
	}
	c.Title = title.String()
}
