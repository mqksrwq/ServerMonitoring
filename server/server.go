package server

import (
	. "ServerMonitoring/task"
	"sync"
	"sync/atomic"
	"time"
)

type Server struct {
	mu        sync.Mutex
	completed int64
	isBusy    atomic.Bool
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(tc <-chan *Task, qc <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case task := <-tc:
			s.isBusy.Store(true)
			s.mu.Lock()
			time.Sleep(task.Duration)
			atomic.AddInt64(&s.completed, 1)
			s.mu.Unlock()

			s.isBusy.Store(false)
		case <-qc:
			return
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
