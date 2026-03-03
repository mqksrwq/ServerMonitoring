package ui

import (
	"ServerMonitoring/server"
	"bufio"
	"fmt"
	"os"
)

// UI - структура интерфейса пользователя
type UI struct {

	// Stack - стек серверов
	Stack *server.Stack

	// scanner - поле для реализации ввода с консоли
	scanner *bufio.Scanner
}

// NewUI - метод, создающий новый ui
func NewUI() *UI {
	return &UI{
		Stack:   server.NewStack(),
		scanner: bufio.NewScanner(os.Stdin),
	}
}

// Run - метод для запуска ui
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
			ui.add()
		default:
			return
		}
	}
}

// add - метод для реализации интерфейса добавления
func (ui *UI) add() {
	ui.Stack.AddTask()
	fmt.Printf("add task\n")
}
