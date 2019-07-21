package main

import (
	"flag"
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

var (
	serial string
)

func init() {
	flag.StringVar(&serial, "serial", "", "serial to connect to uno")
	flag.Parse()
}

func main() {
	fmt.Println("serial", serial)
	if serial == "" {
		panic("serial is required")
	}

	adapter := firmata.NewAdaptor(serial)
	led := gpio.NewLedDriver(adapter, "13")


	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{adapter},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
