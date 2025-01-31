package handlers

import (
	"net/http"
	"path/filepath"
	"strconv"
	"sync"
	"text/template"

	"github.com/ilyas-zal/todo-go/internal/models"
)

// Объявляем переменные для хранения списка задач и мьютекса для синхронизации.
var (
	todos []models.Todo // Слайс для хранения задач.
	mu    sync.Mutex    // Мьютекс для защиты доступа к слайсу todos.
)

// indexHandler обрабатывает GET-запросы на главной странице.
// Он блокирует доступ к списку задач, загружает HTML-шаблон и передает список задач в шаблон для отображения.
func HomeTemplate(w http.ResponseWriter, r *http.Request) {
	mu.Lock()         // Блокируем мьютекс для безопасного доступа к списку задач.
	defer mu.Unlock() // Отпускаем мьютекс после завершения работы функции.
	tmplPath := filepath.Join("frontend", "templates", "index.html")
	// Загружаем HTML-шаблон из файла index.html.
	tmpl := template.Must(template.ParseFiles(tmplPath))
	// Выполняем шаблон, передавая в него список задач и отправляя результат в ответ.
	tmpl.Execute(w, todos)
}

// addHandler обрабатывает POST-запросы для добавления новых задач.
// Если запрос имеет метод POST, он парсит данные формы, добавляет новую задачу в список и перенаправляет пользователя на главную страницу.
func AddTask(w http.ResponseWriter, r *http.Request) {
	// Проверяем, является ли метод запроса POST.
	if r.Method == http.MethodPost {
		// Парсим данные формы из запроса.
		r.ParseForm()
		// Получаем значение поля "task" из формы.
		task := r.FormValue("task")

		mu.Lock() // Блокируем мьютекс для безопасного доступа к списку задач.
		// Добавляем новую задачу в список.
		todos = append(todos, models.Todo{Task: task})
		mu.Unlock() // Отпускаем мьютекс.

		// Перенаправляем пользователя на главную страницу после добавления задачи.
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// completeHandler позволяет изменять статус задачи на выполненную
func CompleteTask(w http.ResponseWriter, r *http.Request) {
	mu.Lock()         // Блокируем мьютекс для безопасного доступа к списку задач.
	defer mu.Unlock() // Отпускаем мьютекс по завершении

	index := r.FormValue("index") // берем индекс задачи
	if index != "" {
		i, err := strconv.Atoi(index) // переводим в инт
		if err == nil && i >= 0 && i < len(todos) {
			// Меняем статус выполнения задачи.
			todos[i].Complete = !todos[i].Complete
		}
	}

	// Перенаправляем пользователя на главную страницу после изменения статуса.
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
