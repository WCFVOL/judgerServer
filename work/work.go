package work

import (
	"awesomeProject/judgerServer/judger"
	"awesomeProject/judgerServer/problem"
)

var taskq = make(chan Task, 2)

// Task judge task or SPJ or upload file
type Task struct {
	TaskID int
	Data   string
}

//DoWork start do work
func DoWork() {
	for {
		select {
		case task := <-taskq:
			if task.TaskID == 1 {
				judger.Handler(task.Data)
			} else if task.TaskID == 2 {
				problem.AddInputOutput(task.Data)
			} else if task.TaskID == 3 {
				// TODO SPJ
			}
		}
	}
}

//AddTask push task to taskq
func AddTask(task Task) {
	taskq <- task
}
