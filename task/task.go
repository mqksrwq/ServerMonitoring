package task

import (
	"math/rand"
	"time"
)

type Task struct {
	Duration time.Duration
}

func NewTask() *Task {
	return &Task{
		Duration: time.Duration(rand.Int63n(int64(10 * time.Second))),
	}
}
