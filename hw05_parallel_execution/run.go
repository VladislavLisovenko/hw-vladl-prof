package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg = sync.WaitGroup{}
	var errCount int32
	limit := make(chan struct{}, n)
	defer close(limit)

	for _, t := range tasks {
		if m > 0 && atomic.LoadInt32(&errCount) >= int32(m) {
			return ErrErrorsLimitExceeded
		}
		limit <- struct{}{}
		wg.Add(1)
		go func(task Task, mm int) {
			defer wg.Done()
			err := task()
			if mm > 0 && err != nil {
				atomic.AddInt32(&errCount, 1)
			}
			<-limit
		}(t, m)
	}

	wg.Wait()

	return nil
}
