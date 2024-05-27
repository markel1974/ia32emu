package core

import (
	"fmt"
	"os"
)

type ALU struct {
	reg *X86Registers
	mem IMemory
}

func NewALU(reg *X86Registers, memory IMemory) *ALU {
	return &ALU{
		reg: reg,
		mem: memory,
	}
}

func (a *ALU) addRM8R8() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r8 := modrm.getR8()
	rm8 := modrm.getRM8()
	modrm.setRM8(rm8 + r8)
}

func (a *ALU) addRM16R16() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r16 := modrm.getR16()
	rm16 := modrm.getRM16()
	modrm.setRM16(rm16 + r16)
}

func (a *ALU) addRM32R32() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r32 := modrm.getR32()
	rm32 := modrm.getRM32()
	modrm.setRM32(rm32 + r32)
}

func (a *ALU) addR8RM8() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r8 := modrm.getR8()
	rm8 := modrm.getRM8()
	modrm.setR8(rm8 + r8)
}

func (a *ALU) addR16RM16() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r16 := modrm.getR16()
	rm16 := modrm.getRM16()
	modrm.setR16(rm16 + r16)
}

func (a *ALU) addR32RM32() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r32 := modrm.getR32()
	rm32 := modrm.getRM32()
	modrm.setR32(rm32 + r32)
}

func (a *ALU) addALImm8() {
	reg := a.reg
	mem := a.mem
	value := mem.GetCode8(1)
	reg.EAX += uint32(value)
	reg.EIP += 2
}

func (a *ALU) addAXImm16() {
	reg := a.reg
	mem := a.mem
	value := mem.GetCode16(1)
	reg.EAX += uint32(value)
	reg.EIP += 3
}

func (a *ALU) addEAXImm32() {
	reg := a.reg
	mem := a.mem
	value := mem.GetCode32(1)
	reg.EAX += value
	reg.EIP += 5
}

func (a *ALU) andRM8R8() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r8 := modrm.getR8()
	rm8 := modrm.getRM8()
	modrm.setRM8(rm8 & r8)
}

func (a *ALU) andRM16R16() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r16 := modrm.getR16()
	rm16 := modrm.getRM16()
	modrm.setRM16(rm16 & r16)
}

func (a *ALU) andRM32R32() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r32 := modrm.getR32()
	rm32 := modrm.getRM32()
	modrm.setRM32(rm32 & r32)
}

func (a *ALU) andR8RM8() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r8 := modrm.getR8()
	rm8 := modrm.getRM8()
	modrm.setR8(rm8 & r8)
}

func (a *ALU) andR16RM16() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r16 := modrm.getR16()
	rm16 := modrm.getRM16()
	modrm.setR16(rm16 & r16)
}

func (a *ALU) andR32RM32() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r32 := modrm.getR32()
	rm32 := modrm.getRM32()
	modrm.setR32(rm32 & r32)
}

func (a *ALU) andALImm8() {
	reg := a.reg
	mem := a.mem
	value := mem.GetCode8(1)
	reg.EAX = reg.EAX & uint32(value)
	reg.EIP += 2
}

func (a *ALU) andAXImm16() {
	reg := a.reg
	mem := a.mem
	value := mem.GetCode16(1)
	reg.EAX = reg.EAX & uint32(value)
	reg.EIP += 3
}

func (a *ALU) andEAXImm32() {
	reg := a.reg
	mem := a.mem
	value := mem.GetCode32(1)
	reg.EAX = reg.EAX & value
	reg.EIP += 5
}

func (a *ALU) orRM8R8() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r8 := modrm.getR8()
	rm8 := modrm.getRM8()
	modrm.setRM8(rm8 | r8)
}

func (a *ALU) orRM16R16() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r16 := modrm.getR16()
	rm16 := modrm.getRM16()
	modrm.setRM16(rm16 | r16)
}

func (a *ALU) orRM32R32() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r32 := modrm.getR32()
	rm32 := modrm.getRM32()
	modrm.setRM32(rm32 | r32)
}

func (a *ALU) orR8RM8() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r8 := modrm.getR8()
	rm8 := modrm.getRM8()
	modrm.setR8(rm8 | r8)
}

func (a *ALU) orR16RM16() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r16 := modrm.getR16()
	rm16 := modrm.getRM16()
	modrm.setR16(rm16 | r16)
}

func (a *ALU) orR32RM32() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r32 := modrm.getR32()
	rm32 := modrm.getRM32()
	modrm.setR32(rm32 | r32)
}

func (a *ALU) orALImm8() {
	reg := a.reg
	mem := a.mem
	value := mem.GetCode8(1)
	reg.EAX = reg.EAX | uint32(value)
	reg.EIP += 2
}

func (a *ALU) orAXImm16() {
	reg := a.reg
	mem := a.mem
	value := mem.GetCode16(1)
	reg.EAX = reg.EAX | uint32(value)
	reg.EIP += 3
}

func (a *ALU) orEAXImm32() {
	reg := a.reg
	mem := a.mem
	value := mem.GetCode32(1)
	reg.EAX = reg.EAX | value
	reg.EIP += 5
}

func (a *ALU) cmpR16RM16() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r16 := modrm.getR16()
	rm16 := modrm.getRM16()
	result := uint32(r16) - uint32(rm16)
	reg.updateEFlagsSub16(r16, rm16, result)
}

func (a *ALU) cmpR32RM32() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r32 := modrm.getR32()
	rm32 := modrm.getRM32()
	result := uint64(r32) - uint64(rm32)
	reg.updateEFlagsSub32(r32, rm32, result)
}

func (a *ALU) cmpALImm8() {
	reg := a.reg
	mem := a.mem
	value := mem.GetCode8(1)
	al := reg.EAX & 0xff
	result := uint64(al) - uint64(value)
	reg.updateEFlagsSub32(al, uint32(value), result)
	reg.EIP += 2
}

func (a *ALU) cmpAXImm16() {
	reg := a.reg
	mem := a.mem
	value := mem.GetCode16(1)
	ax := uint16(reg.EAX & 0xFF)
	result := uint32(ax) - uint32(value)
	reg.updateEFlagsSub16(ax, value, result)
	reg.EIP += 3
}

func (a *ALU) cmpEAXImm32() {
	reg := a.reg
	mem := a.mem
	value := mem.GetCode32(1)
	eax := reg.EAX
	result := uint64(eax) - uint64(value)
	reg.updateEFlagsSub32(eax, value, result)
	reg.EIP += 5
}

func (a *ALU) incR16() {
	reg := a.reg
	mem := a.mem
	index := mem.GetCode8(0) - 0x40
	value := reg.Get16ByIndex(index) + 1
	reg.Set16ByIndex(index, value)
	reg.EIP += 1
}

func (a *ALU) incR32() {
	reg := a.reg
	mem := a.mem
	index := mem.GetCode8(0) - 0x40
	value := reg.GetByIndex(index) + 1
	reg.SetByIndex(index, value)
	reg.EIP += 1
}

func (a *ALU) decR16() {
	reg := a.reg
	mem := a.mem
	index := mem.GetCode8(0) - 0x48
	value := reg.Get16ByIndex(index) - 1
	reg.Set16ByIndex(index, value)
	reg.EIP += 1
}

func (a *ALU) decR32() {
	reg := a.reg
	mem := a.mem
	index := mem.GetCode8(0) - 0x48
	value := reg.GetByIndex(index) - 1
	reg.SetByIndex(index, value)
	reg.EIP += 1
}

func (a *ALU) subRM16R16() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	rm16 := modrm.getRM16()
	r16 := modrm.getR16()
	result := uint32(rm16) - uint32(r16)
	modrm.setRM16(uint16(result))
	reg.updateEFlagsSub16(rm16, r16, result)
}

func (a *ALU) subRM32R32() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	rm32 := modrm.getRM32()
	r32 := modrm.getR32()
	result := uint64(rm32) - uint64(r32)
	modrm.setRM32(uint32(result))
	reg.updateEFlagsSub32(rm32, r32, result)
}

func (a *ALU) subR16RM16() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r16 := modrm.getR16()
	rm16 := modrm.getRM16()
	result := uint32(r16) - uint32(rm16)
	modrm.setR16(uint16(result))
	reg.updateEFlagsSub16(r16, rm16, result)
}

func (a *ALU) subR32RM32() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r32 := modrm.getR32()
	rm32 := modrm.getRM32()
	result := uint64(r32) - uint64(rm32)
	modrm.setR32(uint32(result))
	reg.updateEFlagsSub32(r32, rm32, result)
}

func (a *ALU) subAXImm16() {
	reg := a.reg
	mem := a.mem
	imm16 := mem.GetCode16(1)
	base := reg.EAX
	reg.EAX = base - uint32(imm16)
	reg.EIP += 3
	reg.updateEFlagsSub16(uint16(base), imm16, reg.EAX)
}

func (a *ALU) subEAXImm32() {
	reg := a.reg
	mem := a.mem
	imm32 := mem.GetCode32(1)
	base := reg.EAX
	reg.EAX = reg.EAX - imm32
	reg.EIP += 5
	reg.updateEFlagsSub32(base, imm32, uint64(reg.EAX))
}

func (a *ALU) subRM32Imm32(modrm *ModRM) {
	reg := a.reg
	mem := a.mem
	rm32 := int32(modrm.getRM32())
	imm32 := mem.GetSignCode32(0)
	reg.EIP += 4
	modrm.setRM32(uint32(rm32 - imm32))
}

func (a *ALU) subRM16Imm8(modrm *ModRM) {
	reg := a.reg
	mem := a.mem
	rm16 := modrm.getRM16()
	imm8 := mem.GetSignCode8(0)
	reg.EIP += 1
	result := uint32(rm16) - uint32(imm8)
	modrm.setRM16(uint16(result))
	reg.updateEFlagsSub16(rm16, uint16(imm8), result)
}

func (a *ALU) subRM32Imm8(modrm *ModRM) {
	reg := a.reg
	mem := a.mem
	rm32 := modrm.getRM32()
	imm8 := mem.GetSignCode8(0)
	reg.EIP += 1
	result := uint64(rm32) - uint64(imm8)
	modrm.setRM32(uint32(result))
	reg.updateEFlagsSub32(rm32, uint32(imm8), result)
}

func (a *ALU) subRM8R8() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r8 := modrm.getR8()
	rm8 := modrm.getRM8()
	modrm.setRM8(rm8 - r8)
}

func (a *ALU) subR8RM8() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r8 := modrm.getR8()
	rm8 := modrm.getRM8()
	modrm.setR8(rm8 - r8)
}

func (a *ALU) subALImm8() {
	reg := a.reg
	mem := a.mem
	value := mem.GetCode8(1)
	reg.EAX = reg.EAX - uint32(value)
	reg.EIP += 2
}

func (a *ALU) xorRM16R16() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	rm16 := modrm.getRM16()
	r16 := modrm.getR16()
	modrm.setRM16(rm16 ^ r16)
}

func (a *ALU) xorRM32R32() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	rm32 := modrm.getRM32()
	r32 := modrm.getR32()
	modrm.setRM32(rm32 ^ r32)
}

func (a *ALU) xorR16RM16() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r16 := modrm.getR16()
	rm16 := modrm.getRM16()
	modrm.setR16(r16 ^ rm16)
}

func (a *ALU) xorR32RM32() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r32 := modrm.getR32()
	rm32 := modrm.getRM32()
	modrm.setR32(r32 ^ rm32)
}

func (a *ALU) xorAXImm16() {
	reg := a.reg
	mem := a.mem
	imm16 := mem.GetCode16(1)
	base := reg.EAX
	reg.EAX = base ^ uint32(imm16)
	reg.EIP += 3
}

func (a *ALU) xorEAXImm32() {
	reg := a.reg
	mem := a.mem
	imm32 := mem.GetCode32(1)
	base := reg.EAX
	reg.EAX = base ^ imm32
	reg.EIP += 5
}

func (a *ALU) xorRM8R8() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r8 := modrm.getR8()
	rm8 := modrm.getRM8()
	modrm.setRM8(rm8 ^ r8)
}

func (a *ALU) xorR8RM8() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	r8 := modrm.getR8()
	rm8 := modrm.getRM8()
	modrm.setR8(rm8 ^ r8)
}

func (a *ALU) xorALImm8() {
	reg := a.reg
	mem := a.mem
	value := mem.GetCode8(1)
	reg.EAX = reg.EAX ^ uint32(value)
	reg.EIP += 2
}

func (a *ALU) cmpRM16Imm8(modrm *ModRM) {
	reg := a.reg
	mem := a.mem
	rm16 := modrm.getRM16()
	imm8 := mem.GetSignCode8(0)
	reg.EIP += 1
	result := uint32(rm16) - uint32(imm8)
	reg.updateEFlagsSub16(rm16, uint16(imm8), result)
}

func (a *ALU) cmpRM32Imm8(modrm *ModRM) {
	reg := a.reg
	mem := a.mem
	rm32 := modrm.getRM32()
	imm8 := mem.GetSignCode8(0)
	reg.EIP += 1
	result := uint64(rm32) - uint64(imm8)
	reg.updateEFlagsSub32(rm32, uint32(imm8), result)
}

func (a *ALU) incRM32(modrm *ModRM) {
	value := modrm.getRM32()
	modrm.setRM32(value + 1)
}

func (a *ALU) codeFFb32() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)
	switch modrm.Opcode {
	case 0:
		a.incRM32(&modrm)
	default:
		fmt.Printf("not implemented: FF /%d\n", modrm.Opcode)
		os.Exit(1)
	}
}

func (a *ALU) code83b16() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)

	switch modrm.Opcode {
	case 0:
		//cpu.addRM16Imm8(&modrm) //TODO
	case 1:
		//cpu.orRM16Imm8(&modrm) //TODO
	case 2:
		//cpu.adcRM16Imm8(&modrm) //TODO
	case 3:
		//cpu.sbbRM16Imm8(&modrm) //TODO
	case 4:
		//cpu.andRM16Imm8(&modrm) //TODO
	case 5:
		a.subRM16Imm8(&modrm)
	case 6:
		//cpu.xorRM16Imm8(&modrm) //TODO
	case 7:
		a.cmpRM16Imm8(&modrm)
	default:
		fmt.Printf("not implemented: 0x83 /%d\n", modrm.Opcode)
		os.Exit(1)
	}
}

func (a *ALU) code81b32() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)

	switch modrm.Opcode {
	case 0:
		//cpu.addRM32Imm32(&modrm) //TODO
	case 1:
		//cpu.orRM32Imm32(&modrm) //TODO
	case 2:
		//cpu.adcRM32Imm32(&modrm) //TODO
	case 3:
		//cpu.sbbRM32Imm32(&modrm) //TODO
	case 4:
		//cpu.andRM32Imm32(&modrm) //TODO
	case 5:
		a.subRM32Imm32(&modrm)
	case 6:
		//cpu.xorRM32Imm32(&modrm) //TODO
	case 7:
		//cpu.cmpRM32Imm32(&modrm) //TODO
	default:
		fmt.Printf("not implemented: 0x81 /%d\n", modrm.Opcode)
		os.Exit(1)
	}
}

func (a *ALU) code83b32() {
	reg := a.reg
	reg.EIP += 1
	modrm := NewModRM(a.reg, a.mem)

	switch modrm.Opcode {
	case 0:
		//cpu.addRM32Imm8(&modrm) //TODO
	case 1:
		//cpu.orRM32Imm8(&modrm) //TODO
	case 2:
		//cpu.adcRM32Imm8(&modrm) //TODO
	case 3:
		//cpu.sbbRM32Imm8(&modrm) //TODO
	case 4:
		//cpu.andRM32Imm8(&modrm) //TODO
	case 5:
		a.subRM32Imm8(&modrm)
	case 6:
		//cpu.xorRM32Imm8(&modrm) //TODO
	case 7:
		a.cmpRM32Imm8(&modrm)
	default:
		fmt.Printf("not implemented: 0x83 /%d\n", modrm.Opcode)
		os.Exit(1)
	}
}

func (a *ALU) nop() {
	reg := a.reg
	reg.EIP += 1
}
