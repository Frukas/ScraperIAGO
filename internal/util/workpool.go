package util

import (
	"fmt"
	"sync"
)

type Task interface {
	Run()
}

type workerPool struct {
	WorkNum  int
	TaskChan chan Task
	wg       sync.WaitGroup
}

func WorkFactory(workNum int) *workerPool {
	return &workerPool{
		WorkNum:  workNum,
		TaskChan: make(chan Task, 100),
	}
}

func (wp *workerPool) startWorker() {
	defer wp.wg.Done()
	for task := range wp.TaskChan {
		fmt.Println("TaskRuning")
		defer fmt.Println("Task Ended")
		task.Run()
	}
}

func (wp *workerPool) RunWorker() {

	for i := 0; i < wp.WorkNum; i++ {
		fmt.Println("Workprocess n: ", i)
		wp.wg.Add(1)
		go wp.startWorker()
	}
}

func (wp *workerPool) Wait() {
	wp.wg.Wait()
}
