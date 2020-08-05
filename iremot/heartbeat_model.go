package iremot

import "Iremot-RaspberryPi/device"

var HeartbeatID uint64

type HeartbeatModel struct {
	Id       uint64
	SednTime int64
	PinArr   []device.PinState
}
