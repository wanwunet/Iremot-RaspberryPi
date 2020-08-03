package main

import (
	"Iremot-RaspberryPi/drive"
	"Iremot-RaspberryPi/iremot"
	"fmt"
	"time"
)

func main() {

	drive.Open()
	defer drive.Close()

	iremot.Open()
	defer iremot.Close()

	fmt.Println(drive.GetPinStateAllJson())
	time.Sleep(1000 * time.Second)

}
