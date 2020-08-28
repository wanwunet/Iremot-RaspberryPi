package task

import (
	"Iremot-RaspberryPi/models"
	"fmt"
)

var NowRunTask = make(map[uint64]*models.DeviceTask)

func RegisterTask(t models.DeviceTask) {
	fmt.Println(t)
	NowRunTask[t.Id] = &t
	go t.CallRun(func(Id uint64) {
		delete(NowRunTask, Id)
	})
}
