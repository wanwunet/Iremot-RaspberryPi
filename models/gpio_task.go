package models

type GetGpio struct {
	Pin int8
}

type SetGpio struct {
	Pin   int8
	State uint32
}
