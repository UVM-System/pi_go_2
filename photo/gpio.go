package photo

import (
	"github.com/stianeikeland/go-rpio"
	"pi_go_2/config"
	"sync"
)

var (
	capStartPin rpio.Pin
	capEndPin rpio.Pin
	mutex sync.Mutex
)

// 初始化
func init() {
	if err := rpio.Open(); err != nil {
		panic(err.Error())
	}
	capStartPin = rpio.Pin(config.Config.CapStartPin)
	capEndPin = rpio.Pin(config.Config.CapEndPin)
	capStartPin.Input()
	capEndPin.Input()
}

func GetCapStartPin() rpio.State {
	return capStartPin.Read()
}

func GetCapEndPin() rpio.State {
	return capEndPin.Read()
}

