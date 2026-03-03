package ui

import (
	"ServerMonitoring/server"
	"bufio"
	"fmt"
	"os"
)

type UI struct {
	Stack   *server.Stack
	scanner *bufio.Scanner
}

func NewUI() *UI {
	return &UI{
		Stack:   server.NewStack(),
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (ui *UI) Run() {
	fmt.Println("----- SERVER MONITOR ----")
	someChan := make(chan struct{})
	ui.Stack.StartStack()
	go ui.Stack.Monitoring(ui.Stack.Servers, someChan)
	for {
		ui.scanner.Scan()
		line := ui.scanner.Text()
		switch line {
		case "q":
			close(someChan)
			ui.Stack.StopStack()
			return
		case "1":
			ui.Stack.AddTask()
		default:
			return
		}
	}
}
