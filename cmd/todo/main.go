package main // Объявляем основной пакет, содержащий основную логику приложения.

import (
	"database/sql"
	"fmt"
	"log"
	"net/http" // Импортируем пакет для работы с HTTP.

	"github.com/ilyas-zal/todo-go/internal/handlers" // Импортируем пакет для работы с HTML-шаблонами.

	_ "github.com/go-sql-driver/mysql" // Импортируем пакет для работы с синхронизацией (мьютексы).
)

// main инициализирует маршруты и запускает HTTP-сервер.
// Он связывает обработчики с маршрутами "/" и "/add" и запускает сервер на порту 8080.
func main() {
	fs := http.FileServer(http.Dir("frontend/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlers.HomeTemplate)
	http.HandleFunc("/add", handlers.AddTask)
	http.HandleFunc("/complete", handlers.CompleteTask)

	log.Println("Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v\n", err)
	}
}

func bd() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:4444)/todolist")
	if err != nil {
		fmt.Println("Ошибка подключения к бд")
	}
	defer db.Close()
	fmt.Println("К базе подключились весьма успешно")
}
