package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type DataTwo struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/hello", loggingMiddleware(helloHandlers))
	http.HandleFunc("/data", loggingMiddleware(dataHandlers))
	http.HandleFunc("/goodbye", loggingMiddleware(goodbyeHandler))

	fmt.Println("Сервер запущен на порту 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ошибка при запуске сервера:", err)
	}
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next(w, r)

		duration := time.Since(start)
		log.Printf("%s %s %s", r.Method, r.URL.Path, duration)
	}
}

func helloHandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Привет!"))
}

func dataHandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	var data DataTwo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, "Ошибка при декодировании JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Полученные данные: %+v\n", data)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Данные получены"))
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("До свидания!"))
}
