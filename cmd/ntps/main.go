//package main
//about: btfak.com
//create: 2013-9-25
//update: 2016-08-22

package main

import (
	"fmt"
	"os"
	"strconv"

	"ntp/netapp"
	"ntp/netevent"
	"ntp/sntp"
)

func Main(args []string) {
	if len(args) <= 1 {
		fmt.Println("Usage: ntp port [speed]")
		fmt.Println("	usual default port is 123")
	} else {
		var handler = netapp.GetHandler()
		port, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid port")
		} else {
			if len(args) > 2 {
				speed, err := strconv.ParseInt(args[2], 10, 64)
				if err != nil {
					fmt.Println("Speed error")
					speed = 0
				}
				if speed > 1 {
					sntp.SetSpeedTime(speed - 1)
				}
			}
			netevent.Reactor.ListenUdp(port, handler)
			netevent.Reactor.Run()
		}
	}
}

func main() {
	Main(os.Args)
}
