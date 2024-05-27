package core

type Stack struct {
	reg *X86Registers
	mem IMemory
}

func NewStack(reg *X86Registers, memory IMemory) *Stack {
	return &Stack{
		reg: reg,
		mem: memory,
	}
}

func (s *Stack) PushR32() {
	reg := s.reg
	mem := s.mem
	regIndex := mem.GetCode8(0) - 0x50
	mem.Push32(reg.GetByIndex(regIndex))
	reg.EIP += 1
}

func (s *Stack) Push32Imm32() {
	reg := s.reg
	mem := s.mem
	value := mem.GetCode32(1)
	mem.Push32(value)
	reg.EIP += 5
}

func (s *Stack) Push32Imm8() {
	reg := s.reg
	mem := s.mem
	value := mem.GetCode8(1)
	mem.Push32(uint32(value))
	reg.EIP += 2
}

func (s *Stack) PopR32() {
	reg := s.reg
	mem := s.mem
	regIndex := mem.GetCode8(0) - 0x58
	reg.SetByIndex(regIndex, mem.Pop32())
	reg.EIP += 1
}

func (s *Stack) Push32ES() {
	reg := s.reg
	mem := s.mem
	mem.Push32(uint32(reg.ES))
	reg.EIP += 1
}

func (s *Stack) Pop32ES() {
	reg := s.reg
	mem := s.mem
	reg.ES = uint16(mem.Pop32())
	reg.EIP += 1
}

func (s *Stack) Push32CS() {
	reg := s.reg
	mem := s.mem
	mem.Push32(uint32(reg.CS))
	reg.EIP += 1
}

func (s *Stack) Pop32CS() {
	reg := s.reg
	mem := s.mem
	reg.CS = uint16(mem.Pop32())
	reg.EIP += 1
}

func (s *Stack) Code0Fb32() {
	s.Pop32CS()
}

func (s *Stack) Push32SS() {
	reg := s.reg
	mem := s.mem
	mem.Push32(uint32(reg.SS))
	reg.EIP += 1
}

func (s *Stack) Pop32SS() {
	reg := s.reg
	mem := s.mem
	reg.SS = uint16(mem.Pop32())
	reg.EIP += 1
}

func (s *Stack) Push32DS() {
	reg := s.reg
	mem := s.mem
	mem.Push32(uint32(reg.DS))
	reg.EIP += 1
}

func (s *Stack) Pop32DS() {
	reg := s.reg
	mem := s.mem
	reg.DS = uint16(mem.Pop32())
	reg.EIP += 1
}

// 16 bit instructions

func (s *Stack) Push16ES() {
	reg := s.reg
	mem := s.mem
	mem.Push16(reg.ES)
	reg.EIP += 1
}

func (s *Stack) Pop16ES() {
	reg := s.reg
	mem := s.mem
	reg.ES = mem.Pop16()
	reg.EIP += 1
}

func (s *Stack) Push16CS() {
	reg := s.reg
	mem := s.mem
	mem.Push16(reg.CS)
	reg.EIP += 1
}

func (s *Stack) pop16CS() {
	reg := s.reg
	mem := s.mem
	reg.CS = mem.Pop16()
	reg.EIP += 1
}

func (s *Stack) Code0F16() {
	s.pop16CS()
}

func (s *Stack) Push16SS() {
	reg := s.reg
	mem := s.mem
	mem.Push16(reg.SS)
	reg.EIP += 1
}

func (s *Stack) Pop16SS() {
	reg := s.reg
	mem := s.mem
	reg.SS = mem.Pop16()
	reg.EIP += 1
}

func (s *Stack) Push16DS() {
	reg := s.reg
	mem := s.mem
	mem.Push16(reg.DS)
	reg.EIP += 1
}

func (s *Stack) Pop16DS() {
	reg := s.reg
	mem := s.mem
	reg.DS = mem.Pop16()
	reg.EIP += 1
}

func (s *Stack) Push16Imm16() {
	reg := s.reg
	mem := s.mem
	value := mem.GetCode16(1)
	mem.Push16(value)
	reg.EIP += 3
}

func (s *Stack) Push16Imm8() {
	reg := s.reg
	mem := s.mem
	value := mem.GetCode8(1)
	mem.Push16(uint16(value))
	reg.EIP += 2
}

func (s *Stack) PushR16() {
	reg := s.reg
	mem := s.mem
	regIndex := mem.GetCode8(0) - 0x50
	mem.Push16(reg.Get16ByIndex(regIndex))
	reg.EIP += 1
}

func (s *Stack) PopR16() {
	reg := s.reg
	mem := s.mem
	regIndex := mem.GetCode8(0) - 0x58
	reg.Set16ByIndex(regIndex, mem.Pop16())
	reg.EIP += 1
}
