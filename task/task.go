package task

import (
	"math/rand"
	"time"
)

// Task - структура, имитирующая задачу для сервера
type Task struct {

	// Duration - время выполнения задачи
	Duration time.Duration
}

// NewTask - метод, создающий новую задачу с рандомным временем выполнения
func NewTask() *Task {
	return &Task{
		Duration: time.Duration(rand.Int63n(int64(10 * time.Second))),
	}
}
