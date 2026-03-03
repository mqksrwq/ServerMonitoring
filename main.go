package main

import . "ServerMonitoring/ui"

func main() {
	ui := NewUI()
	ui.Run()
	go ui.Stack.RunMonitor(ui.Stack.Servers)
}
