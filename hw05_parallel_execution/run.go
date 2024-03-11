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
	var errCount int32
	var wg = sync.WaitGroup{}
	limit := make(chan struct{}, n)

	for _, t := range tasks {
		task := t
		if m > 0 && int(errCount) >= m {
			return ErrErrorsLimitExceeded
		}
		limit <- struct{}{}
		wg.Add(1)
		go func() {
			err := task()
			if m > 0 && err != nil {
				atomic.AddInt32(&errCount, 1)
			}
			<-limit
			wg.Done()
		}()
	}

	wg.Wait()

	return nil
}
