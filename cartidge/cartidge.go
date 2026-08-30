package cartidge

import (
	"log/slog"
	"milkboy/cartidge/mbc"
	"os"
	"strings"
)

type Cartidge struct {
	mbc.MBC
	Title    string
	filename string
}

func New(path string) (*Cartidge, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return &Cartidge{}, err
	}

	c := &Cartidge{filename: path}

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

func (c *Cartidge) headerChecksum() byte {
	var checksum byte = 0
	for address := uint16(0x0134); address <= 0x014C; address++ {
		checksum += c.Read(address)
	}
	return checksum
}

func (c *Cartidge) setHeaderChecksum() {
	c.Write(0x014D, c.headerChecksum())
}

func (c *Cartidge) setTitle() {
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
