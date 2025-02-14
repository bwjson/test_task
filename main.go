package main

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"log"
	_ "modernc.org/sqlite" // Подключаем драйвер SQLite
)

//go:embed all:frontend/dist
var assets embed.FS

type Task struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

type TaskManager struct {
	db *sql.DB
}

func NewTaskManager() *TaskManager {
	db, err := sql.Open("sqlite", "tasks.db") // sqlite это драйвер
	if err != nil {
		log.Fatal("Failed to connect to db:", err)
		return nil
	}

	tm := &TaskManager{db: db}

	// Создание таблицы, если она не существует
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        text TEXT NOT NULL,
        completed BOOLEAN NOT NULL DEFAULT 0
    )`)
	if err != nil {
		fmt.Println("Failed to create tables:", err)
	}

	return tm
}

func (tm *TaskManager) AddTask(text string) {
	_, err := tm.db.Exec("INSERT INTO tasks (text, completed) VALUES (?, ?)", text, false)
	if err != nil {
		fmt.Println("Ошибка добавления задачи:", err)
	}
}

func (tm *TaskManager) RemoveTask(id int) {
	_, err := tm.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		fmt.Println("Ошибка удаления задачи:", err)
	}
}

func (tm *TaskManager) ToggleTaskStatus(id int) {
	_, err := tm.db.Exec("UPDATE tasks SET completed = NOT completed WHERE id = ?", id)
	if err != nil {
		fmt.Println("Ошибка изменения статуса задачи:", err)
	}
}

func (tm *TaskManager) GetTasks() []Task {
	rows, err := tm.db.Query("SELECT id, text, completed FROM tasks")
	if err != nil {
		fmt.Println("Ошибка получения задач:", err)
		return nil
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Text, &task.Completed)
		if err != nil {
			fmt.Println("Ошибка чтения задачи:", err)
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:  "TodoApp",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			app.taskManager,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
