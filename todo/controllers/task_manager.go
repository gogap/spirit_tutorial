package controllers

import (
	"errors"
	"github.com/gogap/spirit_tutorial/todo/models"
	"github.com/rs/xid"
	"time"
)

type TaskManager struct {
	tasks map[string]*models.Task
}

var (
	ErrTaskNotExist = errors.New("task not exist")
)

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make(map[string]*models.Task),
	}
}

func (p *TaskManager) NewTask(task models.Task) (id string, err error) {

	task.Id = xid.New().String()
	task.CreateTime = time.Now()
	task.UpdateTime = time.Now()

	p.tasks[task.Id] = &task

	id = task.Id

	return
}

func (p *TaskManager) DeleteTask(taskId string) (err error) {
	if _, exist := p.tasks[taskId]; exist {
		delete(p.tasks, taskId)
	} else {
		err = ErrTaskNotExist
	}
	return
}

func (p *TaskManager) GetTask(taskId string) (task models.Task, err error) {
	if t, exist := p.tasks[taskId]; exist {
		task = *t
	} else {
		err = ErrTaskNotExist
	}
	return
}

func (p *TaskManager) ListTask(userId string) (tasks []models.Task, err error) {
	for _, task := range p.tasks {
		tasks = append(tasks, *task)
	}
	return
}

func (p *TaskManager) FinishTask(taskId string) (err error) {
	if t, exist := p.tasks[taskId]; exist {
		t.IsFinished = true
		t.UpdateTime = time.Now()
	} else {
		err = ErrTaskNotExist
	}
	return
}
