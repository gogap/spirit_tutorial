package component

import (
	"fmt"
	"github.com/gogap/spirit"

	"github.com/gogap/spirit_tutorial/todo/controllers"
	"github.com/gogap/spirit_tutorial/todo/models"
)

var _ spirit.Component = new(TaskComponent)

const (
	taskComponentURN = "urn:spirit_tutorial:component:todo:task"
)

type TaskComponent struct {
	name string
	*controllers.TaskManager
}

func init() {
	spirit.RegisterComponent(taskComponentURN, NewTaskComponent)
}

func NewTaskComponent(name string, options spirit.Map) (component spirit.Component, err error) {
	component = &TaskComponent{
		name:        name,
		TaskManager: controllers.NewTaskManager(),
	}
	return
}

func (p *TaskComponent) Name() string {
	return p.name
}

func (p *TaskComponent) URN() string {
	return taskComponentURN
}

func (p *TaskComponent) Labels() spirit.Labels {
	return spirit.Labels{
		"version": "0.0.1",
	}
}

func (p *TaskComponent) NewTask(payload spirit.Payload) (result interface{}, err error) {
	newTask := models.Task{}

	if err = payload.DataToObject(&newTask); err != nil {
		return
	}

	if result, err = p.TaskManager.NewTask(newTask); err != nil {
		return
	}

	return
}

func (p *TaskComponent) DeleteTask(payload spirit.Payload) (result interface{}, err error) {
	task := models.Task{}

	if err = payload.DataToObject(&task); err != nil {
		return
	}

	if err = p.TaskManager.DeleteTask(task.Id); err != nil {
		return
	}

	return
}

func (p *TaskComponent) GetTask(payload spirit.Payload) (result interface{}, err error) {
	task := models.Task{}

	if err = payload.DataToObject(&task); err != nil {
		return
	}

	if task, err = p.TaskManager.GetTask(task.Id); err != nil {
		return
	}

	result = task

	return
}

func (p *TaskComponent) ListTask(payload spirit.Payload) (result interface{}, err error) {

	if result, err = p.TaskManager.ListTask(); err != nil {
		return
	}

	return
}

func (p *TaskComponent) FinishTask(payload spirit.Payload) (result interface{}, err error) {
	task := models.Task{}

	if err = payload.DataToObject(&task); err != nil {
		return
	}

	if err = p.TaskManager.FinishTask(task.Id); err != nil {
		return
	}

	return
}
