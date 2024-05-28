package core

import (
	"fmt"
	"log"
)

type CPU struct {
	mem      IMemory
	reg      *X86Registers
	debug    bool
	instrSet [0x100]func()
	stack    *Stack
	branch   *Branch
	transfer *Transfer
	io       *IO
	alu      *ALU
}

func NewCPU(reg *X86Registers, mem IMemory, bitMode int, debug bool) *CPU {
	cpu := &CPU{
		mem:      mem,
		reg:      reg,
		debug:    debug,
		stack:    NewStack(reg, mem),
		branch:   NewBranch(reg, mem),
		transfer: NewTransfer(reg, mem),
		io:       NewIO(reg, mem),
		alu:      NewALU(reg, mem),
	}
	if bitMode == 16 {
		cpu.createTable16()
	} else {
		cpu.createTable32()
	}
	return cpu
}

func (cpu *CPU) Init() error {
	//cpu.reg.Init()
	return nil
}

func (cpu *CPU) Exec(code uint8) error {
	if cpu.debug {
		log.Printf("EIP = 0x%X, Opcode = 0x%02X\n", cpu.reg.EIP, code)
	}
	if instr := cpu.instrSet[code]; instr == nil {
		return fmt.Errorf("Not Implemented: 0x%x\n", code)
	} else {
		instr()
	}
	reg := cpu.reg
	if reg.EIP <= cpu.mem.GetAddressBase() {
		return fmt.Errorf("No mapping area [reg.EIP]: 0x%X\n", reg.EIP)
	}
	if uint32(cpu.mem.GetAddressEnd()) <= reg.EIP {
		return fmt.Errorf("No mapping area [mappingEnd]: 0x%X\n", reg.EIP)
	}
	return nil
}

func (cpu *CPU) Dump() {
	cpu.reg.Dump()
}

func (cpu *CPU) createTable16() {
	cpu.instrSet[0x00] = cpu.alu.addRM8R8
	cpu.instrSet[0x01] = cpu.alu.addRM16R16
	cpu.instrSet[0x02] = cpu.alu.addR8RM8
	cpu.instrSet[0x03] = cpu.alu.addR16RM16
	cpu.instrSet[0x04] = cpu.alu.addALImm8
	cpu.instrSet[0x05] = cpu.alu.addAXImm16
	cpu.instrSet[0x06] = cpu.stack.Push16ES
	cpu.instrSet[0x07] = cpu.stack.Pop16ES
	cpu.instrSet[0x08] = cpu.alu.orRM8R8
	cpu.instrSet[0x09] = cpu.alu.orRM16R16
	cpu.instrSet[0x0a] = cpu.alu.orR8RM8
	cpu.instrSet[0x0b] = cpu.alu.orR16RM16
	cpu.instrSet[0x0c] = cpu.alu.orALImm8
	cpu.instrSet[0x0d] = cpu.alu.orAXImm16
	cpu.instrSet[0x0e] = cpu.stack.Push16CS
	cpu.instrSet[0x0f] = cpu.stack.Code0F16

	cpu.instrSet[0x16] = cpu.stack.Push16SS
	cpu.instrSet[0x17] = cpu.stack.Pop16SS
	cpu.instrSet[0x1e] = cpu.stack.Push16DS
	cpu.instrSet[0x1f] = cpu.stack.Pop16DS

	cpu.instrSet[0x21] = cpu.alu.andRM16R16 //TODO VERIFY
	cpu.instrSet[0x22] = cpu.alu.andR8RM8   //TODO VERIFY
	cpu.instrSet[0x23] = cpu.alu.andR16RM16 //TODO VERIFY
	cpu.instrSet[0x24] = cpu.alu.andALImm8  //TODO VERIFY
	cpu.instrSet[0x25] = cpu.alu.andAXImm16 //TODO VERIFY

	cpu.instrSet[0x28] = cpu.alu.subRM8R8 //TODO VERIFY
	cpu.instrSet[0x29] = cpu.alu.subRM16R16
	cpu.instrSet[0x2a] = cpu.alu.subR8RM8 //TODO VERIFY
	cpu.instrSet[0x2b] = cpu.alu.subR16RM16
	cpu.instrSet[0x2c] = cpu.alu.subALImm8 //TODO VERIFY
	cpu.instrSet[0x2d] = cpu.alu.subAXImm16

	cpu.instrSet[0x30] = cpu.alu.xorRM8R8 //TODO VERIFY
	cpu.instrSet[0x31] = cpu.alu.xorRM16R16
	cpu.instrSet[0x32] = cpu.alu.xorR8RM8 //TODO VERIFY
	cpu.instrSet[0x33] = cpu.alu.xorR16RM16
	cpu.instrSet[0x34] = cpu.alu.xorALImm8 //TODO VERIFY
	cpu.instrSet[0x35] = cpu.alu.xorAXImm16

	cpu.instrSet[0x3b] = cpu.alu.cmpR16RM16
	cpu.instrSet[0x3c] = cpu.alu.cmpALImm8
	cpu.instrSet[0x3d] = cpu.alu.cmpAXImm16

	for i := 0; i < 8; i++ {
		cpu.instrSet[0x40+i] = cpu.alu.incR16
	}

	for i := 0; i < 8; i++ {
		cpu.instrSet[0x48+i] = cpu.alu.decR16
	}

	for i := 0; i < 8; i++ {
		cpu.instrSet[0x50+i] = cpu.stack.PushR16
	}

	for i := 0; i < 8; i++ {
		cpu.instrSet[0x58+i] = cpu.stack.PopR16
	}

	cpu.instrSet[0x66] = cpu.overrideOperandTo32
	cpu.instrSet[0x68] = cpu.stack.Push16Imm16
	cpu.instrSet[0x6a] = cpu.stack.Push16Imm8

	cpu.instrSet[0x70] = cpu.branch.JoRel8
	cpu.instrSet[0x71] = cpu.branch.JnoRel8
	cpu.instrSet[0x72] = cpu.branch.JcRel8
	cpu.instrSet[0x73] = cpu.branch.JncRel8
	cpu.instrSet[0x74] = cpu.branch.JzRel8
	cpu.instrSet[0x75] = cpu.branch.JnzRel8
	cpu.instrSet[0x78] = cpu.branch.JsRel8
	cpu.instrSet[0x79] = cpu.branch.JnsRel8
	cpu.instrSet[0x7c] = cpu.branch.JlRel8
	cpu.instrSet[0x7e] = cpu.branch.JleRel8
	cpu.instrSet[0x83] = cpu.alu.code83b16

	cpu.instrSet[0x88] = cpu.transfer.MovRM8R8
	cpu.instrSet[0x89] = cpu.transfer.MovRM16R16
	cpu.instrSet[0x8a] = cpu.transfer.MovR8RM8
	cpu.instrSet[0x8b] = cpu.transfer.MovR16RM16

	cpu.instrSet[0x90] = cpu.alu.nop

	for i := 0; i < 8; i++ {
		cpu.instrSet[0xb0+i] = cpu.transfer.MovR8Imm8
	}

	for i := 0; i < 8; i++ {
		cpu.instrSet[0xb8+i] = cpu.transfer.MovR16Imm16
	}

	cpu.instrSet[0xc3] = cpu.branch.Ret16
	cpu.instrSet[0xc7] = cpu.transfer.MovRM16Imm16
	cpu.instrSet[0xc9] = cpu.branch.Leave16
	/*
		0xd8 - 0xdf: x87 FPU Instructions //TODO
	*/
	cpu.instrSet[0xe8] = cpu.branch.CallRel16
	cpu.instrSet[0xe9] = cpu.branch.JmpRel16
	cpu.instrSet[0xeb] = cpu.branch.JmpRel8
}

func (cpu *CPU) createTable32() {
	cpu.instrSet[0x00] = cpu.alu.addRM8R8
	cpu.instrSet[0x01] = cpu.alu.addRM32R32
	cpu.instrSet[0x02] = cpu.alu.addR8RM8
	cpu.instrSet[0x03] = cpu.alu.addR32RM32
	cpu.instrSet[0x04] = cpu.alu.addALImm8
	cpu.instrSet[0x05] = cpu.alu.addEAXImm32
	cpu.instrSet[0x06] = cpu.stack.Push32ES
	cpu.instrSet[0x07] = cpu.stack.Pop32ES
	cpu.instrSet[0x08] = cpu.alu.orRM8R8
	cpu.instrSet[0x09] = cpu.alu.orRM32R32
	cpu.instrSet[0x0a] = cpu.alu.orR8RM8
	cpu.instrSet[0x0b] = cpu.alu.orR32RM32
	cpu.instrSet[0x0c] = cpu.alu.orALImm8
	cpu.instrSet[0x0d] = cpu.alu.orEAXImm32
	cpu.instrSet[0x0e] = cpu.stack.Push32CS
	cpu.instrSet[0x0f] = cpu.stack.Code0Fb32

	cpu.instrSet[0x16] = cpu.stack.Push32SS
	cpu.instrSet[0x17] = cpu.stack.Pop32SS
	cpu.instrSet[0x1e] = cpu.stack.Push32DS
	cpu.instrSet[0x1f] = cpu.stack.Pop32DS

	cpu.instrSet[0x21] = cpu.alu.andRM32R32  //TODO VERIFY
	cpu.instrSet[0x22] = cpu.alu.andR8RM8    //TODO VERIFY
	cpu.instrSet[0x23] = cpu.alu.andR32RM32  //TODO VERIFY
	cpu.instrSet[0x24] = cpu.alu.andALImm8   //TODO VERIFY
	cpu.instrSet[0x25] = cpu.alu.andEAXImm32 //TODO VERIFY

	cpu.instrSet[0x28] = cpu.alu.subRM8R8 //TODO VERIFY
	cpu.instrSet[0x29] = cpu.alu.subRM32R32
	cpu.instrSet[0x2a] = cpu.alu.subR8RM8 //TODO VERIFY
	cpu.instrSet[0x2b] = cpu.alu.subR32RM32
	cpu.instrSet[0x2c] = cpu.alu.subALImm8 //TODO VERIFY
	cpu.instrSet[0x2d] = cpu.alu.subEAXImm32

	cpu.instrSet[0x30] = cpu.alu.xorRM8R8 //TODO VERIFY
	cpu.instrSet[0x31] = cpu.alu.xorRM32R32
	cpu.instrSet[0x32] = cpu.alu.xorR8RM8 //TODO VERIFY
	cpu.instrSet[0x33] = cpu.alu.xorR32RM32
	cpu.instrSet[0x34] = cpu.alu.xorALImm8 //TODO VERIFY
	cpu.instrSet[0x35] = cpu.alu.xorEAXImm32

	cpu.instrSet[0x3b] = cpu.alu.cmpR32RM32
	cpu.instrSet[0x3c] = cpu.alu.cmpALImm8
	cpu.instrSet[0x3d] = cpu.alu.cmpEAXImm32

	for i := 0; i < 8; i++ {
		cpu.instrSet[0x40+i] = cpu.alu.incR32
	}

	for i := 0; i < 8; i++ {
		cpu.instrSet[0x48+i] = cpu.alu.decR32
	}

	for i := 0; i < 8; i++ {
		cpu.instrSet[0x50+i] = cpu.stack.PushR32
	}

	for i := 0; i < 8; i++ {
		cpu.instrSet[0x58+i] = cpu.stack.PopR32
	}

	cpu.instrSet[0x66] = cpu.overrideOperandTo16
	cpu.instrSet[0x68] = cpu.stack.Push32Imm32
	// cpu.instrSet[0x69] = cpu.imulR32RM32Imm32 //TODO
	cpu.instrSet[0x6a] = cpu.stack.Push32Imm8
	// cpu.instrSet[0x6b] = cpu.imulR32RM32Imm8 //TODO

	cpu.instrSet[0x70] = cpu.branch.JoRel8
	cpu.instrSet[0x71] = cpu.branch.JnoRel8
	cpu.instrSet[0x72] = cpu.branch.JcRel8
	cpu.instrSet[0x73] = cpu.branch.JncRel8
	cpu.instrSet[0x74] = cpu.branch.JzRel8
	cpu.instrSet[0x75] = cpu.branch.JnzRel8
	cpu.instrSet[0x78] = cpu.branch.JsRel8
	cpu.instrSet[0x79] = cpu.branch.JnsRel8
	cpu.instrSet[0x7c] = cpu.branch.JlRel8
	cpu.instrSet[0x7e] = cpu.branch.JleRel8
	cpu.instrSet[0x81] = cpu.alu.code81b32
	cpu.instrSet[0x83] = cpu.alu.code83b32
	// cpu.instrSet[0x84] = cpu.testRM8R8 //TODO
	// cpu.instrSet[0x85] = cpu.testRM32R32 //TODO
	// cpu.instrSet[0x86] = cpu.xchgR8RM8 //TODO
	// cpu.instrSet[0x87] = cpu.xchgR32RM32 //TODO
	cpu.instrSet[0x88] = cpu.transfer.MovRM8R8
	cpu.instrSet[0x89] = cpu.transfer.MovRM32R32
	cpu.instrSet[0x8a] = cpu.transfer.MovR8RM8
	cpu.instrSet[0x8b] = cpu.transfer.MovR32RM32

	cpu.instrSet[0x90] = cpu.alu.nop

	// cpu.instrSet[0xa8] = cpu.testALImm8 //TODO
	// cpu.instrSet[0xa9] = cpu.testEAXImm32 //TODO

	for i := 0; i < 8; i++ {
		cpu.instrSet[0xb0+i] = cpu.transfer.MovR8Imm8
	}

	for i := 0; i < 8; i++ {
		cpu.instrSet[0xb8+i] = cpu.transfer.MovR32Imm32
	}

	cpu.instrSet[0xc3] = cpu.branch.Ret32
	cpu.instrSet[0xc7] = cpu.transfer.MovRM32Imm32
	cpu.instrSet[0xc9] = cpu.branch.Leave32
	/*
		0xd8 - 0xdf: x87 FPU Instructions
	*/
	cpu.instrSet[0xe8] = cpu.branch.CallRel32
	cpu.instrSet[0xe9] = cpu.branch.JmpRel32
	cpu.instrSet[0xeb] = cpu.branch.JmpRel8
	cpu.instrSet[0xec] = cpu.io.InALDX
	cpu.instrSet[0xed] = cpu.io.InEAXDX
	cpu.instrSet[0xee] = cpu.io.OutDXAL
	cpu.instrSet[0xef] = cpu.io.OutDXEAX
	cpu.instrSet[0xff] = cpu.alu.codeFFb32
}

func (cpu *CPU) overrideOperandTo16() {
	reg := cpu.reg
	reg.EIP += 1
	mem := cpu.mem
	code := mem.GetCode8(0)
	if cpu.instrSet[code] == nil {
		log.Fatalf("Not Implemented: 0x%x\n", code)
	}
	cpu.instrSet[code]()
}

func (cpu *CPU) overrideOperandTo32() {
	reg := cpu.reg
	reg.EIP += 1
	mem := cpu.mem
	code := mem.GetCode8(0)
	if cpu.instrSet[code] == nil {
		log.Fatalf("Not Implemented: 0x%x\n", code)
	}
	cpu.instrSet[code]()
}
