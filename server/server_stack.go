package server

import (
	. "ServerMonitoring/task"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Stack - структура стека серверов
type Stack struct {

	// Servers - слайс серверов в стеке
	Servers []*server

	// wg - ожидающая группа
	wg sync.WaitGroup

	// taskChannel - канал для синхронизации задач между горутинами
	taskChannel chan *Task

	// quitChannel - канал для завершения работы серверов
	quitChannel chan struct{}

	// totalComplete - количество выполненных задач стеком
	totalComplete atomic.Int64
}

// NewStack - метод, создающий новый стек серверов
func NewStack() *Stack {
	return &Stack{
		Servers:     make([]*server, 0, 5),
		taskChannel: make(chan *Task, 10),
		quitChannel: make(chan struct{}, 5),
	}
}

// StartStack - метод, запускающий стек серверов
func (s *Stack) StartStack() {
	for i := 0; i < 5; i++ {
		server := newServer()
		s.Servers = append(s.Servers, server)
		s.wg.Add(1)
		go server.startServer(s.taskChannel, s.quitChannel, &s.wg)
	}
}

// StopStack - метод, останавливающий стек и закрывающий каналы
func (s *Stack) StopStack() {
	close(s.taskChannel)
	close(s.quitChannel)
	s.wg.Wait()
}

// AddTask - метод, добавляющий задачу для сервера в очередь
func (s *Stack) AddTask() {
	s.taskChannel <- NewTask()
}

// Monitoring - метод, запускающий мониторинг серверов в стеке
func (s *Stack) Monitoring(servers []*server, qc <-chan struct{}) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			for index, server := range servers {
				fmt.Printf("\033[2K%s\n", server.toString(index))
			}
			fmt.Printf("-----\n")
		case <-qc:
			return
		}
	}
}
