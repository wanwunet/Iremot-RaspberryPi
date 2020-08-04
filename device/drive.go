package device

import "Iremot-RaspberryPi/device/rpio"

func Open() {
	rpio.Open()
}

func Close() {
	rpio.Close()
}
