package main

import (
	"Iremot-RaspberryPi/conf"
	"Iremot-RaspberryPi/device"
	"Iremot-RaspberryPi/iremot"
	_ "Iremot-RaspberryPi/www/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	fmt.Println(conf.Serial())
	fmt.Println(conf.Model())
	fmt.Println(conf.Ether())

	device.Open()
	defer device.Close()

	iremot.Open()
	defer iremot.Close()

	beego.Run()
}
