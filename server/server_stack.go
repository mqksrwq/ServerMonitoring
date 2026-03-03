package server

import (
	. "ServerMonitoring/task"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Stack struct {
	Servers       []*Server
	taskChannel   chan *Task
	wg            sync.WaitGroup
	quitChannel   chan struct{}
	totalComplete atomic.Int64
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Start() {

	for i := 0; i < 5; i++ {
		server := NewServer()
		s.Servers = append(s.Servers, server)
		s.wg.Add(1)
		go server.Run(s.taskChannel, s.quitChannel, &s.wg)
	}
}

func (s *Stack) Stop() {
	close(s.taskChannel)
	close(s.quitChannel)
	s.wg.Wait()
}

func (s *Stack) AddTask() {
	s.taskChannel <- NewTask()
}

func (s *Stack) RunMonitor(servers []*Server) {
	for {
		select {
		case <-time.After(2 * time.Second):
			for _, server := range servers {
				fmt.Println(server)
			}
		case <-s.quitChannel:
			return
		}
	}
}
