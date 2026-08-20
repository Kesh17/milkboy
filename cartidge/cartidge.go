package cartidge

import (
	"os"
)

type Cartidge struct {
	ROM []byte
}

func New(path string) (*Cartidge, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return &Cartidge{}, err
	}

	return &Cartidge{ROM: data}, nil
}

func (c *Cartidge) Read(addr uint16) byte {
	return c.ROM[addr]
}
