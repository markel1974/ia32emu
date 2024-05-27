package core

import (
	"fmt"
	"os"
)

type ModRM struct {
	reg *X86Registers
	mem IMemory

	Mod      uint8
	Rm       uint8
	Opcode   uint8
	RegIndex uint8
	Sib      uint8
	Disp8    int8
	Disp32   uint32
}

func NewModRM(reg *X86Registers, mem IMemory) ModRM {
	modrm := ModRM{reg: reg, mem: mem}
	code := mem.GetCode8(0)

	modrm.Mod = (code & 0xc0) >> 6
	modrm.Opcode = (code & 0x38) >> 3
	modrm.RegIndex = modrm.Opcode
	modrm.Rm = code & 0x7

	reg.EIP += 1

	if modrm.Mod != 3 && modrm.Rm == 4 {
		modrm.Sib = mem.GetCode8(0)
		reg.EIP += 1
	}
	if (modrm.Mod == 0 && modrm.Rm == 5) || modrm.Mod == 2 {
		modrm.Disp32 = mem.GetCode32(0)
		reg.EIP += 4
	} else if modrm.Mod == 1 {
		modrm.Disp8 = mem.GetSignCode8(0)
		modrm.Disp32 = uint32(modrm.Disp8)
		reg.EIP += 1
	}
	return modrm
}

func (modrm *ModRM) calcAddress() uint32 {
	if modrm.Mod == 0 {
		if modrm.Rm == 4 {
			fmt.Println("not implemented ModRM mod = 0, rm = 4")
			os.Exit(0)
		}
		if modrm.Rm == 5 {
			return modrm.Disp32
		}
		result := modrm.reg.GetByIndex(modrm.Rm)
		return result
	}

	if modrm.Mod == 1 {
		if modrm.Rm == 4 {
			fmt.Println("not implemented ModRM mod = 2, rm = 4")
			os.Exit(0)
		}
		result := modrm.reg.GetByIndex(modrm.Rm) + modrm.Disp32
		return result
	}

	fmt.Println("not implemented ModRM mod", modrm.Mod)
	os.Exit(0)
	return 0
}

func (modrm *ModRM) setRM8(value uint8) {
	if modrm.Mod == 3 {
		reg := modrm.reg
		reg.Set8ByIndex(modrm.Rm, value)
	} else {
		mem := modrm.mem
		address := modrm.calcAddress()
		mem.Write8(address, value)
	}
}

func (modrm *ModRM) setRM16(value uint16) {
	if modrm.Mod == 3 {
		reg := modrm.reg
		reg.Set16ByIndex(modrm.Rm, value)
	} else {
		mem := modrm.mem
		address := modrm.calcAddress()
		mem.Write16(address, value)
	}
}

func (modrm *ModRM) setRM32(value uint32) {
	if modrm.Mod == 3 {
		reg := modrm.reg
		reg.SetByIndex(modrm.Rm, value)
	} else {
		mem := modrm.mem
		address := modrm.calcAddress()
		mem.Write32(address, value)
	}
}

func (modrm *ModRM) getRM8() (result uint8) {
	if modrm.Mod == 3 {
		reg := modrm.reg
		result = reg.Get8ByIndex(modrm.Rm)
	} else {
		mem := modrm.mem
		address := modrm.calcAddress()
		result = mem.Read8(address)
	}
	return result
}

func (modrm *ModRM) getRM16() (result uint16) {
	if modrm.Mod == 3 {
		reg := modrm.reg
		result = reg.Get16ByIndex(modrm.Rm)
	} else {
		mem := modrm.mem
		address := modrm.calcAddress()
		result = mem.Read16(address)
	}
	return result
}

func (modrm *ModRM) getRM32() (result uint32) {
	if modrm.Mod == 3 {
		reg := modrm.reg
		result = reg.GetByIndex(modrm.Rm)
	} else {
		mem := modrm.mem
		address := modrm.calcAddress()
		result = mem.Read32(address)
	}
	return result
}

func (modrm *ModRM) setR8(value uint8) {
	reg := modrm.reg
	reg.Set8ByIndex(modrm.RegIndex, value)
}

func (modrm *ModRM) setR16(value uint16) {
	reg := modrm.reg
	reg.Set16ByIndex(modrm.RegIndex, value)
}

func (modrm *ModRM) setR32(value uint32) {
	reg := modrm.reg
	reg.SetByIndex(modrm.RegIndex, value)
}

func (modrm *ModRM) getR8() uint8 {
	reg := modrm.reg
	return reg.Get8ByIndex(modrm.RegIndex)
}

func (modrm *ModRM) getR16() uint16 {
	reg := modrm.reg
	return reg.Get16ByIndex(modrm.RegIndex)
}

func (modrm *ModRM) getR32() uint32 {
	reg := modrm.reg
	return reg.GetByIndex(modrm.RegIndex)
}
