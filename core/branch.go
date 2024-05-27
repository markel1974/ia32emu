package core

type Branch struct {
	reg *X86Registers
	mem IMemory
}

func NewBranch(reg *X86Registers, memory IMemory) *Branch {
	return &Branch{
		reg: reg,
		mem: memory,
	}
}

func (b *Branch) JoRel8() {
	reg := b.reg
	diff := uint32(2)
	if reg.IsOF() {
		mem := b.mem
		diff += uint32(mem.GetCode8(1))
	}
	reg.EIP += diff
}

func (b *Branch) JnoRel8() {
	reg := b.reg
	diff := uint32(2)
	if !reg.IsOF() {
		mem := b.mem
		diff += uint32(mem.GetCode8(1))
	}
	reg.EIP += diff
}

func (b *Branch) JcRel8() {
	reg := b.reg
	diff := uint32(2)
	if reg.IsCF() {
		mem := b.mem
		diff += uint32(mem.GetCode8(1))
	}
	reg.EIP += diff
}

func (b *Branch) JncRel8() {
	reg := b.reg
	diff := uint32(2)
	if !reg.IsCF() {
		mem := b.mem
		diff += uint32(mem.GetCode8(1))
	}
	reg.EIP += diff
}

func (b *Branch) JzRel8() {
	reg := b.reg
	diff := uint32(2)
	if reg.IsZF() {
		mem := b.mem
		diff += uint32(mem.GetCode8(1))
	}
	reg.EIP += diff
}

func (b *Branch) JnzRel8() {
	reg := b.reg
	diff := uint32(2)
	if !reg.IsZF() {
		mem := b.mem
		diff += uint32(mem.GetCode8(1))
	}
	reg.EIP += diff
}

func (b *Branch) JnzRel16() {
	reg := b.reg
	diff := uint16(4)
	if !reg.IsZF() {
		mem := b.mem
		diff += mem.GetCode16(2)
	}
	reg.EIP = uint32(uint16(reg.EIP) + diff)
}

func (b *Branch) JnzRel32() {
	reg := b.reg
	diff := uint32(6)
	if !reg.IsZF() {
		mem := b.mem
		diff += mem.GetCode32(2)
	}
	reg.EIP += diff
}

func (b *Branch) JsRel8() {
	reg := b.reg
	diff := uint32(2)
	if reg.IsSF() {
		mem := b.mem
		diff += uint32(mem.GetCode8(1))
	}
	reg.EIP += diff
}

func (b *Branch) JnsRel8() {
	reg := b.reg
	diff := uint32(2)
	if !reg.IsSF() {
		mem := b.mem
		diff += uint32(mem.GetCode8(1))
	}
	reg.EIP += diff
}

func (b *Branch) JlRel8() {
	reg := b.reg
	diff := uint32(2)
	if reg.IsSF() != reg.IsOF() {
		mem := b.mem
		diff += uint32(mem.GetCode8(1))
	}
	reg.EIP += diff
}

func (b *Branch) JleRel8() {
	reg := b.reg
	diff := uint32(2)
	if reg.IsZF() || (reg.IsSF() != reg.IsOF()) {
		mem := b.mem
		diff += uint32(mem.GetCode8(1))
	}
	reg.EIP += diff
}

func (b *Branch) JmpRel32() {
	reg := b.reg
	mem := b.mem
	diff := mem.GetCode32(1)
	reg.EIP += diff + 5
}

func (b *Branch) JmpRel16() {
	reg := b.reg
	mem := b.mem
	diff := mem.GetCode16(1)
	reg.EIP = uint32(uint16(reg.EIP) + diff + 3)
}

func (b *Branch) JmpRel8() {
	reg := b.reg
	mem := b.mem
	diff := mem.GetCode8(1)
	reg.EIP += uint32(diff + 2)
}

func (b *Branch) Ret32() {
	reg := b.reg
	mem := b.mem
	reg.EIP = mem.Pop32()
}

func (b *Branch) Leave32() {
	reg := b.reg
	mem := b.mem
	ebp := reg.EBP
	reg.ESP = ebp
	reg.EBP = mem.Pop32()
	reg.EIP += 1
}

func (b *Branch) CallRel32() {
	reg := b.reg
	mem := b.mem
	diff := mem.GetSignCode32(1)
	mem.Push32(reg.EIP + 5)
	reg.EIP += uint32(diff + 5)
}

func (b *Branch) Ret16() {
	reg := b.reg
	mem := b.mem
	reg.EIP = uint32(mem.Pop16())
}

func (b *Branch) Leave16() {
	reg := b.reg
	mem := b.mem
	ebp := reg.EBP
	reg.ESP = ebp
	reg.EBP = uint32(mem.Pop16())
	reg.EIP += 1
}

func (b *Branch) CallRel16() {
	reg := b.reg
	mem := b.mem
	diff := mem.GetSignCode16(1)
	mem.Push16(uint16(reg.EIP + 3))
	reg.EIP += uint32(diff) + 3
}
