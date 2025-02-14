import './style.css';
import './app.css';

import { GetTasks, AddTask, RemoveTask, ToggleTaskStatus } from '../wailsjs/go/main/App';

document.querySelector('#app').innerHTML = `
    <h1>Todo List</h1>
    <input id="taskInput" type="text" placeholder="Новая задача">
    <button id="addTask">Добавить</button>
    <div id="taskList"></div>
`;

const taskInput = document.getElementById("taskInput");
const addTaskButton = document.getElementById("addTask");
const taskList = document.getElementById("taskList");


function loadTasks() {
    GetTasks().then(tasks => {
        taskList.innerHTML = "";
        tasks.forEach(task => {
            const taskElement = document.createElement("div");
            taskElement.classList.add("task");
            if (task.completed) taskElement.classList.add("completed");

            taskElement.innerHTML = `
                <span>${task.text}</span>
                <button onclick="toggleTask(${task.id})">${task.completed ? "✅" : "❌"}</button>
                <button onclick="removeTask(${task.id})">🗑</button>
            `;
            taskList.appendChild(taskElement);
        });
    });
}


addTaskButton.addEventListener("click", () => {
    const text = taskInput.value.trim();
    if (text) {
        AddTask(text).then(() => {
            taskInput.value = "";
            loadTasks();
        });
    }
});


window.toggleTask = function (id) {
    ToggleTaskStatus(id).then(() => loadTasks());
};


window.removeTask = function (id) {
    RemoveTask(id).then(() => loadTasks());
};


loadTasks();
