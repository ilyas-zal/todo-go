package main // Объявляем основной пакет, содержащий основную логику приложения.

import (
	// "database/sql"
	"html/template" // Импортируем пакет для работы с HTML-шаблонами.
	"net/http"      // Импортируем пакет для работы с HTTP.
	"strconv"
	"sync"
	//_ "github.com/go-sql-driver/mysql" // Импортируем пакет для работы с синхронизацией (мьютексы).
)

// Todo представляет структуру задачи.
type Todo struct {
	Task     string // Поле Task хранит текст задачи.
	Complete bool   // Проверка, выполнена ли задача
}

// Объявляем переменные для хранения списка задач и мьютекса для синхронизации.
var (
	todos []Todo     // Слайс для хранения задач.
	mu    sync.Mutex // Мьютекс для защиты доступа к слайсу todos.
)

// main инициализирует маршруты и запускает HTTP-сервер.
// Он связывает обработчики с маршрутами "/" и "/add" и запускает сервер на порту 8080.
func main() {
	// Обрабатываем запросы на главной странице, связывая их с обработчиком indexHandler.
	http.HandleFunc("/", indexHandler)
	// Обрабатываем запросы на добавление задач, связывая их с обработчиком addHandler.
	http.HandleFunc("/add", addHandler)

	http.HandleFunc("/complete", completeHandler)

	// Запускаем HTTP-сервер на порту 8080.
	http.ListenAndServe(":8080", nil)
}

// indexHandler обрабатывает GET-запросы на главной странице.
// Он блокирует доступ к списку задач, загружает HTML-шаблон и передает список задач в шаблон для отображения.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()         // Блокируем мьютекс для безопасного доступа к списку задач.
	defer mu.Unlock() // Отпускаем мьютекс после завершения работы функции.

	// Загружаем HTML-шаблон из файла index.html.
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	// Выполняем шаблон, передавая в него список задач и отправляя результат в ответ.
	tmpl.Execute(w, todos)
}

// addHandler обрабатывает POST-запросы для добавления новых задач.
// Если запрос имеет метод POST, он парсит данные формы, добавляет новую задачу в список и перенаправляет пользователя на главную страницу.
func addHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, является ли метод запроса POST.
	if r.Method == http.MethodPost {
		// Парсим данные формы из запроса.
		r.ParseForm()
		// Получаем значение поля "task" из формы.
		task := r.FormValue("task")

		mu.Lock() // Блокируем мьютекс для безопасного доступа к списку задач.
		// Добавляем новую задачу в список.
		todos = append(todos, Todo{Task: task})
		mu.Unlock() // Отпускаем мьютекс.

		// Перенаправляем пользователя на главную страницу после добавления задачи.
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// completeHandler позволяет изменять статус задачи на выполненную
func completeHandler(w http.ResponseWriter, r *http.Request) {
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
