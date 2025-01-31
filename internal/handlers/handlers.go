package handlers

import (
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/ilyas-zal/todo-go/internal/todo"
)

// Объявляем переменные для хранения списка задач и мьютекса для синхронизации.
var todoService = todo.NewTodoService()

// indexHandler обрабатывает GET-запросы на главной странице.
// Он блокирует доступ к списку задач, загружает HTML-шаблон и передает список задач в шаблон для отображения.
func HomeTemplate(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("frontend", "templates", "index.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	tmpl.Execute(w, todoService.GetTasks())
}

// addHandler обрабатывает POST-запросы для добавления новых задач.
// Если запрос имеет метод POST, он парсит данные формы, добавляет новую задачу в список и перенаправляет пользователя на главную страницу.
func AddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		task := r.FormValue("task")
		todoService.AddTask(task)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// completeHandler позволяет изменять статус задачи на выполненную
func CompleteTask(w http.ResponseWriter, r *http.Request) {
	index := r.FormValue("index")
	if index != "" {
		i, _ := strconv.Atoi(index)
		todoService.CompleteTask(i)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
