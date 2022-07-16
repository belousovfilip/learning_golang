package hw05parallelexecution

import (
	"errors"
	"sync"
)

var maxErrors int

var countErrors int

var mx = sync.Mutex{}

var wg = sync.WaitGroup{}

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

var ErrErrorsNegativeNumberThreads = errors.New("negative number threads")

type Task func() error

func Run(tasks []Task, n, m int) error {
	maxErrors = m
	countErrors = 0
	if n < 1 {
		return ErrErrorsNegativeNumberThreads
	}
	taskChunks := splitTasksToChunks(tasks, n)
	for _, taskChunk := range taskChunks {
		wg.Add(1)
		go func(taskChunk []Task) {
			defer wg.Done()
			for _, task := range taskChunk {
				if cannotHandleTasks() {
					break
				}
				handleTask(task)
			}
		}(taskChunk)
	}
	wg.Wait()
	if countErrors > maxErrors {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func handleTask(task Task) {
	if err := task(); err != nil {
		incrementCountErrors()
	}
}

func cannotHandleTasks() bool {
	mx.Lock()
	defer mx.Unlock()
	return countErrors > maxErrors
}

func incrementCountErrors() {
	mx.Lock()
	countErrors++
	mx.Unlock()
}

func splitTasksToChunks(tasks []Task, size int) [][]Task {
	chunks := make([][]Task, size)
	lenTasks := len(tasks)
	lenChunk := lenTasks / size
	low := 0
	high := lenChunk
	for i := 0; i < size; i++ {
		if i+1 == size {
			high = lenTasks
		}
		chunks[i] = append(chunks[i], tasks[low:high]...)
		low += lenChunk
		high += lenChunk
	}
	return chunks
}
