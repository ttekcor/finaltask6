package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi/v5"
)

type Service struct {
	log *log.Logger
	server http.Server
	router *http.ServeMux
	
}

func NewServer() *http.Server {
	// Создаем логгер
	logger := log.New(log.Writer(), "SERVER: ", log.LstdFlags)
	
	// Создаем роутер
	r := chi.NewRouter()
	r.Get("/", handlers.HandlerHTML)
	r.Post("/upload", handlers.Upload)
	
	// Создаем экземпляр http.Server с указанными настройками
	server := &http.Server{
		Addr:         ":8080",           // Порт 8080
		Handler:      r,                 // HTTP-роутер
		ErrorLog:     logger,            // Логгер
		ReadTimeout:  5 * time.Second,   // Таймаут для чтения - 5 секунд
		WriteTimeout: 10 * time.Second,  // Таймаут для записи - 10 секунд
		IdleTimeout:  15 * time.Second,  // Таймаут ожидания - 15 секунд
	}
	
	return server
}

func main() {
	// Создаем сервер
	_ = NewServer()
	// Здесь можно запустить сервер
}