package models

type GetGpio struct {
	Pin int8
}

type SetGpio struct {
	Pin   int8
	State uint32
}

type PwmGpio struct {
	Pin      int8
	Freq     int
	DutyLen  uint32
	CycleLen uint32
}

type DataFrameGpio struct {
	Pin     int8
	State   []uint32
	Delayed int64
}

type DefGpio struct {
	Pin int8
}
