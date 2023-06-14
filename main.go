package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"io"
)

func main() {
	// Устанавливаем обработчики эндпоинтов
	http.HandleFunc("/get_html", getHTMLHandler)
	http.HandleFunc("/say_hello", sayHelloHandler)
	http.HandleFunc("/say_current_time", sayCurrentTimeHandler)
	http.HandleFunc("/get_info", getInfoHandler)

	// Запускаем сервер на порту 9999
	port := 9999
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Server has been started on port %d", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func getHTMLHandler(w http.ResponseWriter, r *http.Request) {
	// Открываем файл index.html
	file, err := http.Dir(".").Open("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Копируем содержимое файла в ResponseWriter
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func sayCurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем текущее время по МСК
	currentTime := time.Now().UTC().Add(time.Hour * 3)

	// Отправляем время в формате строки
	fmt.Fprint(w, currentTime.Format("15:04:05"))
}

func getInfoHandler (w http.ResponseWriter, r *http.Request) {
	// открываем фаил info.html
	file, err := http.Dir(".").Open("info.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		
	}
	defer file.Close()

	// Копируем содержимое файла в ResponsWriter

	_, err = io.Copy(w, file)
	if  err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
} 
