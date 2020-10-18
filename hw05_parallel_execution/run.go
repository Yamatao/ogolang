package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
var ErrWrongArguments = errors.New("wrong input parameters")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks.
func Run(tasks []Task, N int, M int) error {
	if M <= 0 || N <= 0 {
		return ErrWrongArguments
	}
	if len(tasks) == 0 {
		return nil
	}

	// create N goroutines
	wg := sync.WaitGroup{}
	stopAll := make(chan interface{})
	worker := func(_ int, tasks <-chan Task, stop <-chan interface{}, errs chan<- error) {
		defer wg.Done()
		for active := true; active; {
			select {
			case task, ok := <-tasks:
				if !ok {
					return
				}
				err := task()
				if err != nil {
					errs <- err
					return
				}
			case <-stop:
				return
			}
		}
	}

	tasksChan, errsChan := make(chan Task), make(chan error)
	for i := 0; i < N; i++ {
		wg.Add(1)
		go worker(i, tasksChan, stopAll, errsChan)
	}

	// send tasks
	numErrors := 0
	defer wg.Wait()
	defer close(tasksChan)
	for taskIdx := 0; taskIdx < len(tasks); {
		select {
		case tasksChan <- tasks[taskIdx]:
			taskIdx++
		case <-errsChan:
			numErrors++
			if numErrors == M || numErrors == N {
				close(stopAll)
				return ErrErrorsLimitExceeded
			}
		}
	}
	return nil
}
