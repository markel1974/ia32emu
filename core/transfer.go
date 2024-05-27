package core

type Transfer struct {
	reg *X86Registers
	mem IMemory
}

func NewTransfer(reg *X86Registers, memory IMemory) *Transfer {
	return &Transfer{
		reg: reg,
		mem: memory,
	}
}

func (t *Transfer) MovRM8R8() {
	reg := t.reg
	reg.EIP += 1
	modrm := NewModRM(t.reg, t.mem)
	r8 := modrm.getR8()
	modrm.setRM8(r8)
}

func (t *Transfer) MovR8RM8() {
	reg := t.reg
	reg.EIP += 1
	modrm := NewModRM(t.reg, t.mem)
	rm8 := modrm.getRM8()
	modrm.setR8(rm8)
}

func (t *Transfer) MovRM16R16() {
	reg := t.reg
	reg.EIP += 1
	modrm := NewModRM(t.reg, t.mem)
	r16 := modrm.getR16()
	modrm.setRM16(r16)
}

func (t *Transfer) MovR16RM16() {
	reg := t.reg
	reg.EIP += 1
	modrm := NewModRM(t.reg, t.mem)
	rm16 := modrm.getRM16()
	modrm.setR16(rm16)
}

func (t *Transfer) MovRM32R32() {
	reg := t.reg
	reg.EIP += 1
	modrm := NewModRM(t.reg, t.mem)
	r32 := modrm.getR32()
	modrm.setRM32(r32)
}

func (t *Transfer) MovR32RM32() {
	reg := t.reg
	reg.EIP += 1
	modrm := NewModRM(t.reg, t.mem)
	rm32 := modrm.getRM32()
	modrm.setR32(rm32)
}

func (t *Transfer) MovR8Imm8() {
	mem := t.mem
	regIndex := mem.GetCode8(0) - 0xb0
	imm8 := mem.GetCode8(1)
	reg := t.reg
	reg.Set8ByIndex(regIndex, imm8)
	reg.EIP += 2
}

func (t *Transfer) MovR16Imm16() {
	mem := t.mem
	regIndex := mem.GetCode8(0) - 0xb8
	imm16 := mem.GetCode16(1)
	reg := t.reg
	reg.Set16ByIndex(regIndex, imm16)
	reg.EIP += 3
}

func (t *Transfer) MovRM16Imm16() {
	reg := t.reg
	mem := t.mem
	reg.EIP += 1
	modrm := NewModRM(t.reg, t.mem)
	imm16 := mem.GetCode16(0)
	reg.EIP += 2
	modrm.setRM16(imm16)
}

func (t *Transfer) MovR32Imm32() {
	mem := t.mem
	regIndex := mem.GetCode8(0) - 0xb8
	imm32 := mem.GetCode32(1)
	reg := t.reg
	reg.SetByIndex(regIndex, imm32)
	reg.EIP += 5
}

func (t *Transfer) MovRM32Imm32() {
	reg := t.reg
	mem := t.mem
	reg.EIP += 1
	modrm := NewModRM(t.reg, t.mem)
	imm32 := mem.GetCode32(0)
	reg.EIP += 4
	modrm.setRM32(imm32)
}
