package models

type HeartbeatModel struct {
	Id       uint64
	SednTime int64
	PinArr   []PinState
}

type PinState struct {
	Pin   int8
	Mod   uint32
	State uint32
}
