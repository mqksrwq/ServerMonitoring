package server

import (
	. "ServerMonitoring/task"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Stack struct {
	Servers       []*server
	wg            sync.WaitGroup
	taskChannel   chan *Task
	quitChannel   chan struct{}
	totalComplete atomic.Int64
	PrintMutex    sync.Mutex
}

func NewStack() *Stack {
	return &Stack{
		Servers:     make([]*server, 0, 5),
		taskChannel: make(chan *Task, 10),
		quitChannel: make(chan struct{}, 5),
	}
}

func (s *Stack) StartStack() {

	for i := 0; i < 5; i++ {
		server := newServer()
		s.Servers = append(s.Servers, server)
		s.wg.Add(1)
		go server.startServer(s.taskChannel, s.quitChannel, &s.wg)
	}
}

func (s *Stack) StopStack() {
	close(s.taskChannel)
	close(s.quitChannel)
	s.wg.Wait()
}

func (s *Stack) AddTask() {
	s.taskChannel <- NewTask()
}

func (s *Stack) Monitoring(servers []*server, qc <-chan struct{}) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			s.PrintMutex.Lock()
			fmt.Print("\033[H\033[2J")
			for index, server := range servers {
				fmt.Printf("\033[2K%s\n", server.toString(index))
			}
			fmt.Printf("-----\n")
			s.PrintMutex.Unlock()
		case <-qc:
			return
		}
	}
}
