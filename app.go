package main

import (
	"context"
)

type App struct {
	ctx         context.Context
	taskManager *TaskManager
}

func NewApp() *App {
	return &App{
		taskManager: NewTaskManager(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetTasks() []Task {
	return a.taskManager.GetTasks()
}

func (a *App) AddTask(text string) {
	a.taskManager.AddTask(text)
}

func (a *App) RemoveTask(id int) {
	a.taskManager.RemoveTask(id)
}

func (a *App) ToggleTaskStatus(id int) {
	a.taskManager.ToggleTaskStatus(id)
}
