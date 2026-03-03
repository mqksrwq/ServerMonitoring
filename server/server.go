package server

import (
	. "ServerMonitoring/task"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// server - структура условного сервера
type server struct {

	// mu - блокировщик записи
	mu sync.Mutex

	// completed - количество выполненных сервером задач
	completed int64

	// isBusy - булевое значение, занят/не занят
	isBusy atomic.Bool
}

// newServer - метод, создающий новый сервер
func newServer() *server {
	return &server{}
}

// startServer - метод, запускающий определенный сервер
func (s *server) startServer(tc <-chan *Task, qc <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case task, ok := <-tc:
			if !ok {
				return
			}
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

// toString - метод, конвертирующий структуру сервера в удобочитаемую строку
func (s *server) toString(i int) string {
	return fmt.Sprintf("server %d\tcompleted %d\t"+
		"isBusy %t\t", i, s.completed, s.isBusy.Load())
}
