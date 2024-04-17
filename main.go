package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"net/http"
)

func main() {
	// Создаем сервер
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		// Отдаем текст "Hello, web!"
		fmt.Fprintf(w, "Hello, web!")
	})

	// Запускаем сервер
	http.ListenAndServe(":8080", nil)
}
