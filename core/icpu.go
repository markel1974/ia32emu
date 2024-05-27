package core

type ICPU interface {
	Init() error
	Exec(code uint8) error
	Dump()
}
