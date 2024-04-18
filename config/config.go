package config

import (
	"fmt"
	"net/http"
)

func Hello() {
	// Создаем сервер
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Отдаем текст "Hello, web!"
		fmt.Fprintf(w, "Hello, web!")
	})

	// Запускаем сервер
	http.ListenAndServe(":8080", nil)
}
