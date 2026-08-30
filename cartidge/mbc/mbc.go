package mbc

type MBC interface {
	Read(addr uint16) byte
	Write(addr uint16, data byte)
	WriteRam(addr uint16, data byte)
}
