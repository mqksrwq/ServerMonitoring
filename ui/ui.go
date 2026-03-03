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
	go ui.Stack.Start()
}
