package core

type Emulator struct {
	cpu ICPU
	mem IMemory
	// TODO:  devices
}

func NewEmulator(bitMode int, baseAddress uint32, stackAddress uint32, ram []byte, debug bool) (*Emulator, error) {
	reg := NewIA32registers(baseAddress, stackAddress, debug)
	mem := NewMemory(reg, ram, baseAddress, debug)
	cpu := NewCPU(reg, mem, bitMode, debug)
	emu := &Emulator{cpu: cpu, mem: mem}
	return emu, nil
}

func (emu *Emulator) SetRam(ram []byte) {
	emu.mem.SetRam(ram)
}

func (emu *Emulator) Run() error {
	if err := emu.cpu.Init(); err != nil {
		return err
	}

	for {
		code := emu.mem.GetCode8(0)
		if err := emu.cpu.Exec(code); err != nil {
			return err
		}
	}
}

func (emu *Emulator) Dump() {
	emu.cpu.Dump()
}
