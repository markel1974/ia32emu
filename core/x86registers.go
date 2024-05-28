package core

import (
	"fmt"
	"log"
	"reflect"
)

type X86Registers struct {
	// GPR
	EAX uint32
	ECX uint32
	EDX uint32
	EBX uint32
	ESP uint32
	EBP uint32
	ESI uint32
	EDI uint32

	// Instruction Register
	EIP uint32

	// Segment Registers
	CS uint16
	DS uint16
	SS uint16
	ES uint16
	FS uint16
	GS uint16

	// FLAGS Register
	EFlags uint32
	// MMX registers (MM0 through MM7)
	MM0 uint64
	MM1 uint64
	MM2 uint64
	MM3 uint64
	MM4 uint64
	MM5 uint64
	MM6 uint64
	MM7 uint64
	// TODO: XMM registers (XMM0 through XMM15) and the MXCSR register
	// uint128 doesn't existed

	// Control Registers
	CR0 uint64
	CR1 uint64
	CR2 uint64
	CR3 uint64
	CR4 uint64
	CR5 uint64
	CR6 uint64
	CR7 uint64

	//baseAddress  uint32
	//stackAddress uint32
	debug bool
}

func NewIA32registers(baseAddress uint32, stackAddress uint32, debug bool) *X86Registers {
	r := &X86Registers{
		//baseAddress:  baseAddress,
		//stackAddress: stackAddress,
		debug: debug,
	}
	r.init(baseAddress, stackAddress)
	return r
}

func (r *X86Registers) init(baseAddress uint32, stackAddress uint32) {
	r.EAX = 0
	r.ECX = 0
	r.EDX = 0
	r.EBX = 0
	r.ESP = stackAddress
	r.EBP = 0
	r.ESI = 0
	r.EDI = 0

	r.EIP = baseAddress

	r.CS = 0
	r.DS = 0
	r.SS = 0
	r.ES = 0
	r.FS = 0
	r.GS = 0

	// Reserved 1st bit, it's always 1 in EFlags.
	r.EFlags = 2

	r.MM0 = 0
	r.MM1 = 0
	r.MM2 = 0
	r.MM3 = 0
	r.MM4 = 0
	r.MM5 = 0
	r.MM6 = 0
	r.MM7 = 0

	r.CR0 = 0
	r.CR1 = 0
	r.CR2 = 0
	r.CR3 = 0
	r.CR4 = 0
	r.CR5 = 0
	r.CR6 = 0
	r.CR7 = 0
}

func (r *X86Registers) registerIndex(index uint8) *uint32 {
	switch index {
	case 0:
		return &r.EAX
	case 1:
		return &r.ECX
	case 2:
		return &r.EDX
	case 3:
		return &r.EBX
	case 4:
		return &r.ESP
	case 5:
		return &r.EBP
	case 6:
		return &r.ESI
	case 7:
		return &r.EDI
	}
	log.Fatal("UNDEFINED index", index)
	return &r.EAX
}

func (r *X86Registers) Dump() {
	v := reflect.ValueOf(r).Elem()
	t := v.Type()

	fmt.Println("==================== registers ====================")
	for i := 0; i < 24; i++ {
		registerName := t.Field(i).Name
		registerValue := v.Field(i).Interface()

		switch registerName {
		case "EFlags":
			fmt.Printf("%02d: %s = 0x%X (%032b)\n",
				i+1, registerName, registerValue, registerValue)
		default:
			fmt.Printf("%02d: %s = 0x%X\n",
				i+1, registerName, registerValue)
		}
	}
}

func (r *X86Registers) GetByIndex(index uint8) uint32 {
	value32 := r.registerIndex(index)
	return *value32
}

func (r *X86Registers) Get16ByIndex(index uint8) uint16 {
	value32 := r.registerIndex(index)
	value16 := uint16(*value32 & 0xFFFF)
	return value16
}

func (r *X86Registers) Get8ByIndex(index uint8) uint8 {
	value32 := r.registerIndex(index)
	value8 := uint8(*value32 & 0xFF)
	return value8
}

func (r *X86Registers) SetByIndex(index uint8, value uint32) {
	value32 := r.registerIndex(index)
	*value32 = value
}

func (r *X86Registers) Set16ByIndex(index uint8, value uint16) {
	value32 := r.registerIndex(index)
	value16 := (*value32 & 0xFFFF0000) + uint32(value)
	*value32 = value16
}

func (r *X86Registers) Set8ByIndex(index uint8, value uint8) {
	if index <= 3 {
		v := r.registerIndex(index)
		value8 := (*v & 0xFFFF00FF) + (uint32(value) << 8)
		*v = value8
	} else {
		value32 := r.registerIndex(index - 4)
		value8 := (*value32 & 0xFFFF00FF) + (uint32(value) << 8)
		*value32 = value8
	}
}

func (r *X86Registers) updateEFlagsAdd16(v1 uint16, v2 uint16, result uint32) {
	sign1 := int(v1 >> 15)
	sign2 := int(v2 >> 15)
	signR := int((result >> 15) & 1)

	if (result >> 16) == 1 {
		r.SetCF()
	} else {
		r.RemoveCF()
	}

	if result == 0 {
		r.SetZF()
	} else {
		r.RemoveZF()
	}

	if signR == 1 {
		r.SetSF()
	} else {
		r.RemoveSF()
	}

	if (sign1^sign2 == 0) && (sign1^signR == 1) {
		r.SetOF()
	} else {
		r.RemoveOF()
	}
}

func (r *X86Registers) updateEFlagsAdd32(v1 uint32, v2 uint32, result uint64) {
	sign1 := int(v1 >> 31)
	sign2 := int(v2 >> 31)
	signR := int((result >> 31) & 1)

	if (result >> 32) == 1 {
		r.SetCF()
	} else {
		r.RemoveCF()
	}

	if result == 0 {
		r.SetZF()
	} else {
		r.RemoveZF()
	}

	if signR == 1 {
		r.SetSF()
	} else {
		r.RemoveSF()
	}

	if (sign1^sign2 == 0) && (sign1^signR == 1) {
		r.SetOF()
	} else {
		r.RemoveOF()
	}
}

func (r *X86Registers) updateEFlagsOr16(result uint16) {
	r.RemoveCF()

	if result == 0 {
		r.SetZF()
	} else {
		r.RemoveZF()
	}

	if ((result >> 15) & 1) == 1 {
		r.SetSF()
	} else {
		r.RemoveSF()
	}

	r.RemoveOF()
}

func (r *X86Registers) updateEFlagsOr32(result uint32) {
	r.RemoveCF()

	if result == 0 {
		r.SetZF()
	} else {
		r.RemoveZF()
	}

	if ((result >> 31) & 1) == 1 {
		r.SetSF()
	} else {
		r.RemoveSF()
	}

	r.RemoveOF()
}

func (r *X86Registers) updateEFlagsAnd16(result uint16) {
	r.RemoveCF()

	if result == 0 {
		r.SetZF()
	} else {
		r.RemoveZF()
	}

	if ((result >> 15) & 1) == 1 {
		r.SetSF()
	} else {
		r.RemoveSF()
	}

	r.RemoveOF()
}

func (r *X86Registers) updateEFlagsAnd32(result uint32) {
	r.RemoveCF()

	if result == 0 {
		r.SetZF()
	} else {
		r.RemoveZF()
	}

	if ((result >> 31) & 1) == 1 {
		r.SetSF()
	} else {
		r.RemoveSF()
	}

	r.RemoveOF()
}

func (r *X86Registers) updateEFlagsSub16(v1 uint16, v2 uint16, result uint32) {
	sign1 := int(v1 >> 15)
	sign2 := int(v2 >> 15)
	signR := int((result >> 15) & 1)

	if (result >> 16) == 1 {
		r.SetCF()
	} else {
		r.RemoveCF()
	}

	if result == 0 {
		r.SetZF()
	} else {
		r.RemoveZF()
	}

	if signR == 1 {
		r.SetSF()
	} else {
		r.RemoveSF()
	}

	if (sign1^sign2 == 1) && (sign1^signR == 1) {
		r.SetOF()
	} else {
		r.RemoveOF()
	}
}

func (r *X86Registers) updateEFlagsSub32(v1 uint32, v2 uint32, result uint64) {
	sign1 := int(v1 >> 31)
	sign2 := int(v2 >> 31)
	signR := int((result >> 31) & 1)

	if (result >> 32) == 1 {
		r.SetCF()
	} else {
		r.RemoveCF()
	}

	if result == 0 {
		r.SetZF()
	} else {
		r.RemoveZF()
	}

	if signR == 1 {
		r.SetSF()
	} else {
		r.RemoveSF()
	}

	if (sign1^sign2 == 1) && (sign1^signR == 1) {
		r.SetOF()
	} else {
		r.RemoveOF()
	}
}

func (r *X86Registers) updateEFlagsMul16(result uint64) {
	msb := result >> 16
	if msb > 0 {
		r.SetCF()
		r.SetOF()
	} else {
		r.RemoveCF()
		r.RemoveOF()
	}
}

func (r *X86Registers) updateEFlagsMul32(result uint64) {
	msb := result >> 32
	if msb > 0 {
		r.SetCF()
		r.SetOF()
	} else {
		r.RemoveCF()
		r.RemoveOF()
	}
}

func (r *X86Registers) updateEFlagsImul16(result uint64) {
	msb := result >> 16
	if int(msb) != -1 {
		r.SetCF()
		r.SetOF()
	} else {
		r.RemoveCF()
		r.RemoveOF()
	}
}

func (r *X86Registers) updateEFlagsImul32(result uint64) {
	msb := result >> 32
	if int(msb) != -1 {
		r.SetCF()
		r.SetOF()
	} else {
		r.RemoveCF()
		r.RemoveOF()
	}
}

// IsCF FLAGS Register Carry Flag (0 bit)
func (r *X86Registers) IsCF() bool {
	return (r.EFlags & 1) != 0
}

func (r *X86Registers) SetCF() {
	r.EFlags = r.EFlags | 1
}

func (r *X86Registers) RemoveCF() {
	mask := ^1
	r.EFlags &= uint32(mask)
}

// IsPF Parity Flag (2bit)
func (r *X86Registers) IsPF() bool {
	return (r.EFlags & 4) != 0
}

func (r *X86Registers) SetPF() {
	r.EFlags = r.EFlags | 4
}

func (r *X86Registers) RemovePF() {
	mask := ^4
	r.EFlags &= uint32(mask)
}

// IsAF Adjust Flag (4bit)
func (r *X86Registers) IsAF() bool {
	return (r.EFlags & 16) != 0
}

func (r *X86Registers) SetAF() {
	r.EFlags = r.EFlags | 16
}

func (r *X86Registers) RemoveAF() {
	mask := ^16
	r.EFlags &= uint32(mask)
}

// IsZF Zero Flag (6bit)
func (r *X86Registers) IsZF() bool {
	return (r.EFlags & 64) != 0
}

func (r *X86Registers) SetZF() {
	r.EFlags = r.EFlags | 64
}

func (r *X86Registers) RemoveZF() {
	mask := ^64
	r.EFlags &= uint32(mask)
}

// IsSF Sign Flag (7bit)
func (r *X86Registers) IsSF() bool {
	return (r.EFlags & 128) != 0
}

func (r *X86Registers) SetSF() {
	r.EFlags = r.EFlags | 128
}

func (r *X86Registers) RemoveSF() {
	mask := ^128
	r.EFlags &= uint32(mask)
}

// IsTF Trap Flag (8bit)
func (r *X86Registers) IsTF() bool {
	return (r.EFlags & 256) != 0
}

func (r *X86Registers) SetTF() {
	r.EFlags = r.EFlags | 256
}

func (r *X86Registers) RemoveTF() {
	mask := ^256
	r.EFlags &= uint32(mask)

}

// IsEF Interrupt Enable Flag (9bit)
func (r *X86Registers) IsEF() bool {
	return (r.EFlags & 512) != 0
}

func (r *X86Registers) SetIF() {
	r.EFlags = r.EFlags | 512
}

func (r *X86Registers) RemoveIF() {
	mask := ^512
	r.EFlags &= uint32(mask)
}

// IsDF Direction Flag (10bit)
func (r *X86Registers) IsDF() bool {
	return (r.EFlags & 1024) != 0
}

func (r *X86Registers) SetDF() {
	r.EFlags = r.EFlags | 1024
}

func (r *X86Registers) RemoveDF() {
	mask := ^1024
	r.EFlags &= uint32(mask)
}

// IsOF Overflow Flag (11bit)
func (r *X86Registers) IsOF() bool {
	return (r.EFlags & 2048) != 0
}

func (r *X86Registers) SetOF() {
	r.EFlags = r.EFlags | 2048
}

func (r *X86Registers) RemoveOF() {
	mask := ^2048
	r.EFlags &= uint32(mask)
}

// IsIOPL I/O Privilege Level Field (12-13bit)
func (r *X86Registers) IsIOPL() bool {
	return (r.EFlags & 4096) != 0
}

func (r *X86Registers) SetIOPL() {
	r.EFlags = r.EFlags | 4096 // TODO: fix
}

func (r *X86Registers) RemoveIOPL() {
	mask := ^4096 // TODO: fix
	r.EFlags &= uint32(mask)
}

// IsNT Nested Task Flag (14bit)
func (r *X86Registers) IsNT() bool {
	return (r.EFlags & 16384) != 0
}

func (r *X86Registers) SetNT() {
	r.EFlags = r.EFlags | 16384
}

func (r *X86Registers) RemoveNT() {
	mask := ^16384
	r.EFlags &= uint32(mask)
}

// IsRF Resume Flag (16bit)
func (r *X86Registers) IsRF() bool {
	return (r.EFlags & 65536) != 0
}

func (r *X86Registers) SetRF() {
	r.EFlags = r.EFlags | 65536
}

func (r *X86Registers) RemoveRF() {
	mask := ^65536
	r.EFlags &= uint32(mask)
}

// IsVM Virtual x86 Mode Flag (17bit)
func (r *X86Registers) IsVM() bool {
	return (r.EFlags & 131072) != 0
}

func (r *X86Registers) SetVM() {
	r.EFlags = r.EFlags | 131072
}

func (r *X86Registers) RemoveVM() {
	mask := ^131072
	r.EFlags &= uint32(mask)
}

// IsAC Alignment Check Flag (18bit)
func (r *X86Registers) IsAC() bool {
	return (r.EFlags & 262144) != 0
}

func (r *X86Registers) SetAC() {
	r.EFlags = r.EFlags | 262144
}

func (r *X86Registers) RemoveAC() {
	mask := ^262144
	r.EFlags &= uint32(mask)
}

// IsVIF Virtual Interrupt Flag (19bit)
func (r *X86Registers) IsVIF() bool {
	return (r.EFlags & 524288) != 0
}

func (r *X86Registers) SetVIF() {
	r.EFlags = r.EFlags | 524288
}

func (r *X86Registers) RemoveVIF() {
	mask := ^524288
	r.EFlags &= uint32(mask)
}

// IsVIP Virtual Interrupt Pending Flag (20bit)
func (r *X86Registers) IsVIP() bool {
	return (r.EFlags & 1048576) != 0
}

func (r *X86Registers) SetVIP() {
	r.EFlags = r.EFlags | 1048576
}

func (r *X86Registers) RemoveVIP() {
	mask := ^1048576
	r.EFlags &= uint32(mask)
}

// IsID Identification Flag (21bit)
func (r *X86Registers) IsID() bool {
	return (r.EFlags & 2097152) != 0
}

func (r *X86Registers) SetID() {
	r.EFlags = r.EFlags | 2097152
}

func (r *X86Registers) RemoveID() {
	mask := ^2097152
	r.EFlags &= uint32(mask)
}

func (r *X86Registers) CheckParity(v uint8) bool {
	var p uint8 = 1
	var i uint8
	for i = 0; i < 8; i++ {
		p ^= (v >> i) & 1
	}
	return p == 1
}
