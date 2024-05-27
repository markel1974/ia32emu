package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type IO struct {
	reg *X86Registers
	mem IMemory
}

func NewIO(reg *X86Registers, memory IMemory) *IO {
	return &IO{
		reg: reg,
		mem: memory,
	}
}

func (i *IO) InALDX() {
	reg := i.reg
	address := uint16(reg.EDX & 0xffff)
	value := ioIn8(address)
	AH := reg.EAX & 0xff00
	reg.EAX = AH + uint32(value)
	reg.EIP += 1
}

func (i *IO) InEAXDX() {
	reg := i.reg
	address := uint16(reg.EDX & 0xffff)
	value := ioIn32(address)
	reg.EAX = value
	reg.EIP += 1
}

func (i *IO) OutDXAL() {
	reg := i.reg
	address := uint16(reg.EDX & 0xffff)
	AL := uint8(reg.EAX & 0xff)
	ioOut8(address, AL)
	reg.EIP += 1
}

func (i *IO) OutDXEAX() {
	reg := i.reg
	address := uint16(reg.EDX & 0xffff)
	ioOut32(address, reg.EAX)
	reg.EIP += 1
}

func ioIn8(address uint16) uint8 {
	fmt.Println("ioIn8 input ...")
	switch address {
	case 0x03c7: // Palette Address(Read Mode) on VGA
		break
	case 0x03c9: // Palette Data on VGA
		break
	case 0x03cc: // Miscellaneous Output Register on VGA
		break
	case 0x03f8: // COM1
		var input []byte = make([]byte, 1)
		os.Stdin.Read(input)
		return uint8(input[0])
		break
	}
	return 0
}

func ioIn32(address uint16) uint32 {
	fmt.Println("ioIn32 input ...")
	switch address {
	case 0x03c7: // Palette Address(Read Mode) on VGA
		break
	case 0x03c9: // Palette Data on VGA
		break
	case 0x03cc: // Miscellaneous Output Register on VGA
		break
	case 0x03f8: // COM1
		var input = make([]byte, 4)
		os.Stdin.Read(input)
		var i uint32
		buf := bytes.NewReader(input)
		binary.Read(buf, binary.LittleEndian, &i)
		return i
		break
	}
	return 0
}

func ioOut8(address uint16, ascii uint8) {
	switch address {
	case 0x03c2: // Miscellaneous Output Register on VGA
		break
	case 0x03c8: // Palette Address(Write Mode) on VGA
		break
	case 0x03c9: // Palette Data on VGA
		break
	case 0x03f8: // COM1
		fmt.Println(string(ascii))
		break
	}
}

func ioOut32(address uint16, ascii uint32) {
	switch address {
	case 0x03c2: // Miscellaneous Output Register on VGA
		break
	case 0x03c8: // Palette Address(Write Mode) on VGA
		break
	case 0x03c9: // Palette Data on VGA
		break
	case 0x03f8: // COM1
		fmt.Println(string(ascii))
		break
	}
}
