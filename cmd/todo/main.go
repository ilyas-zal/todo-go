package main // Объявляем основной пакет, содержащий основную логику приложения.

import (
	"database/sql"
	"fmt"
	"net/http" // Импортируем пакет для работы с HTTP.

	"github.com/ilyas-zal/todo-go/internal/handlers" // Импортируем пакет для работы с HTML-шаблонами.

	_ "github.com/go-sql-driver/mysql" // Импортируем пакет для работы с синхронизацией (мьютексы).
)

// main инициализирует маршруты и запускает HTTP-сервер.
// Он связывает обработчики с маршрутами "/" и "/add" и запускает сервер на порту 8080.
func main() {
	//bd()
	// Обрабатываем запросы на главной странице, связывая их с обработчиком indexHandler.
	http.HandleFunc("/", handlers.HomeTemplate)
	// Обрабатываем запросы на добавление задач, связывая их с обработчиком addHandler.
	http.HandleFunc("/add", handlers.AddTask)

	http.HandleFunc("/complete", handlers.CompleteTask)

	// Запускаем HTTP-сервер на порту 8080.
	http.ListenAndServe(":8080", nil)

}

func bd() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:4444)/todolist")
	if err != nil {
		fmt.Println("Ошибка подключения к бд")
	}
	defer db.Close()
	fmt.Println("К базе подключились весьма успешно")
}
