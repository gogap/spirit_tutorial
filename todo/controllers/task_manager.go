package controllers

import (
	"time"

	"github.com/gogap/errors"
	"github.com/rs/xid"

	"github.com/gogap/spirit_tutorial/todo/models"
)

type TaskManager struct {
	tasks map[string]*models.Task
}

var (
	ErrTaskNotExist = errors.TN(TodoErrorNamespace, 50404, "task not exist, id: {{.id}}")
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
		err = ErrTaskNotExist.New(errors.Params{"id": taskId})
	}
	return
}

func (p *TaskManager) GetTask(taskId string) (task models.Task, err error) {
	if t, exist := p.tasks[taskId]; exist {
		task = *t
	} else {
		err = ErrTaskNotExist.New(errors.Params{"id": taskId})
	}
	return
}

func (p *TaskManager) ListTask() (tasks []models.Task, err error) {
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
		err = ErrTaskNotExist.New(errors.Params{"id": taskId})
	}
	return
}
