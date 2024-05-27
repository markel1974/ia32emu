package core

type IMemory interface {
	SetRam(ram []uint8)
	RamSize() uint64
	GetAddressBase() uint32
	GetAddressEnd() uint64
	Read(address uint32) byte
	Read8(address uint32) uint8
	Read16(address uint32) uint16
	Read32(address uint32) uint32
	Write(address uint32, value byte)
	Write8(address uint32, value uint8)
	Write16(address uint32, value uint16)
	Write32(address uint32, value uint32)
	GetCode8(offset int) uint8
	GetSignCode8(offset int) int8
	GetCode16(offset int) uint16
	GetSignCode16(offset int) int16
	GetCode32(offset int) uint32
	GetSignCode32(offset int) int32
	Push16(value uint16)
	Push32(value uint32)
	Pop16() uint16
	Pop32() uint32
}
