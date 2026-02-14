package main

import (
	"FinalProject_G92/network"
	"fmt"
	"net"
)

func wvManager(worldviewCh chan [network.N]network.Call) {

	var wv [network.N]network.Call
	var floor int
	var dir string

	//taking keyboard input for tests
	for {
		fmt.Print("Floor and direction (e.g. '2 u'): ")
		fmt.Scan(&floor, &dir)
		if floor >= 0 && floor < network.N {
			switch dir {
			case "u":
				wv[floor].Up = !wv[floor].Up
			case "d":
				wv[floor].Down = !wv[floor].Down
			default:
				fmt.Println("Use 'u' or 'd'")
				continue
			}
			worldviewCh <- wv
		}
	}
}

func main() {
	ip := net.ParseIP("127.0.0.1")

	heartbeatCh := make(chan network.Heartbeat)
	worldviewCh := make(chan [network.N]network.Call)

	go network.Listener(heartbeatCh, ip)
	go network.Heart(worldviewCh, ip)

	//worldview
	go wvManager(worldviewCh)

	// Print heartbeats as they arrive
	for hb := range heartbeatCh {
		network.PrintWorldView(hb.Worldview)
	}
}
