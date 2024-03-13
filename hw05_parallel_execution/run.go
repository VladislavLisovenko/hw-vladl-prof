package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	workersCount := n
	if workersCount > len(tasks) {
		workersCount = len(tasks)
	}

	job := make(chan Task)

	var wg sync.WaitGroup
	var errCount int32

	for i := 0; i < workersCount; i++ {
		wg.Add(1)

		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for t := range job {
				err := t()
				if m > 0 && err != nil {
					atomic.AddInt32(&errCount, 1)
				}
			}
		}(&wg)
	}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for _, t := range tasks {
			if m > 0 && atomic.LoadInt32(&errCount) >= int32(m) {
				break
			}
			job <- t
		}
		close(job)
	}(&wg)

	wg.Wait()

	if m > 0 && errCount >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
