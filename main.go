package main

import (
	"Iremot-RaspberryPi/device"
	"Iremot-RaspberryPi/iremot"
	_ "Iremot-RaspberryPi/www/routers"
	"github.com/astaxie/beego"
)

func main() {

	device.Open()
	defer device.Close()

	iremot.Open()
	defer iremot.Close()

	beego.Run()
}
