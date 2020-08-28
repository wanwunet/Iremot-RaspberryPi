package models

type DeviceTask struct {
	Id        uint64
	Name      string
	Loop      int64
	TaskFrame []DataFrame
}

type DataFrame struct {
	PinSta  []DataFramePinSta
	Delayed uint64
}

type DataFramePinSta struct {
	Pin int
	Sta int
}
