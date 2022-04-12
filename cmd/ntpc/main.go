package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"ntp/sntp"
)

func Main(args []string) {
	if len(args) <= 1 {
		fmt.Println("Usage: ADDRESS[:PORT] [ADDRESS[:PORT] ...]")
		fmt.Println("	default port is 123")
		// 0.beevik-ntp.pool.ntp.org 0.fr.pool.ntp.org 0.europe.pool.ntp.org 0.uk.pool.ntp.org
	} else {
		for i := 1; i < len(args); i++ {
			var host string
			var port int
			if strings.Contains(args[i], ":") {
				host = strings.Split(args[i], ":")[0]
				if sport, err := strconv.Atoi(strings.Split(args[i], ":")[1]); err == nil {
					port = sport
				} else {
					port = 123
				}
			} else {
				host = args[i]
				port = 123
			}
			res, err := sntp.Client(host, port)
			if err != nil {
				fmt.Println(args[i], " => Unable to get time")
			} else {
				fmt.Printf("%s => %v\n", args[i], res)
			}
		}
	}
}

func main() {
	Main(os.Args)
}
