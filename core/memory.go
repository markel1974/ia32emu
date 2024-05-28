package core

type Memory struct {
	addressBase uint32
	addressEnd  uint64
	reg         *X86Registers
	//TODO must be cpu aligned!! 32/64 bit
	ram   []byte
	debug bool
}

func NewMemory(reg *X86Registers, ram []byte, addressBase uint32, debug bool) *Memory {
	return &Memory{
		reg:         reg,
		addressBase: addressBase,
		ram:         ram,
		addressEnd:  uint64(addressBase) + uint64(len(ram)),
		debug:       debug,
	}
}

func (mem *Memory) GetAddressBase() uint32 {
	return mem.addressBase
}

func (mem *Memory) GetAddressEnd() uint64 {
	return mem.addressEnd
}

func (mem *Memory) SetRam(ram []uint8) {
	mem.ram = ram
	mem.addressEnd = uint64(mem.addressBase) + uint64(len(mem.ram))
}

func (mem *Memory) RamSize() uint64 {
	return uint64(len(mem.ram))
}

func (mem *Memory) Read(address uint32) byte {
	index := address - mem.addressBase
	return mem.ram[index]
}

func (mem *Memory) Read8(address uint32) uint8 {
	return mem.Read(address)
}

func (mem *Memory) Read16(address uint32) uint16 {
	var ret uint16
	for i := uint32(0); i < 2; i++ {
		ret |= uint16(mem.Read(address+i)) << (8 * i)
	}
	return ret
}

func (mem *Memory) Read32(address uint32) uint32 {
	var ret uint32
	for i := uint32(0); i < 4; i++ {
		ret |= uint32(mem.Read(address+i)) << (8 * i)
	}
	return ret
}

func (mem *Memory) Write(address uint32, value byte) {
	index := address - mem.addressBase
	mem.ram[index] = value
}

func (mem *Memory) Write8(address uint32, value uint8) {
	mem.Write(address, value)
}

func (mem *Memory) Write16(address uint32, value uint16) {
	for i := 0; i < 2; i++ {
		mem.Write(address+uint32(i), byte(value>>(uint(i)*8)))
	}
}

func (mem *Memory) Write32(address uint32, value uint32) {
	for i := 0; i < 4; i++ {
		mem.Write(address+uint32(i), byte(value>>(uint(i)*8)))
	}
}

func (mem *Memory) GetCode8(offset int) uint8 {
	reg := mem.reg
	return mem.Read(reg.EIP + uint32(offset))
}

func (mem *Memory) GetSignCode8(offset int) int8 {
	reg := mem.reg
	return int8(mem.Read(reg.EIP + uint32(offset)))
}

func (mem *Memory) GetCode16(offset int) uint16 {
	var i uint
	var ret uint16
	for i = 0; i < 2; i++ {
		ret |= uint16(mem.GetCode8(offset+int(i))) << (i * 8)
	}
	return ret
}

func (mem *Memory) GetSignCode16(offset int) int16 {
	return int16(mem.GetCode16(offset))
}

func (mem *Memory) GetCode32(offset int) uint32 {
	var i uint
	var ret uint32
	for i = 0; i < 4; i++ {
		ret |= uint32(mem.GetCode8(offset+int(i))) << (i * 8)
	}
	return ret
}

func (mem *Memory) GetSignCode32(offset int) int32 {
	return int32(mem.GetCode32(offset))
}

func (mem *Memory) Push16(value uint16) {
	reg := mem.reg
	address := reg.ESP - 2
	reg.ESP = address
	mem.Write16(address, value)
}

func (mem *Memory) Push32(value uint32) {
	reg := mem.reg
	address := reg.ESP - 4
	reg.ESP = address
	mem.Write32(address, value)
}

func (mem *Memory) Pop16() (ret uint16) {
	reg := mem.reg
	address := reg.ESP
	value := mem.Read16(address)
	reg.ESP = address + 2
	return value
}

func (mem *Memory) Pop32() (ret uint32) {
	reg := mem.reg
	address := reg.ESP
	value := mem.Read32(address)
	reg.ESP = address + 4
	return value
}
