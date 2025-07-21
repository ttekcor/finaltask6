package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(log.Writer(), "MAIN: ", log.LstdFlags)
	server := server.NewServer()
	logger.Printf("Сервер запускается на порту 8080...")


	if err := server.ListenAndServe(); err != nil {
		logger.Fatal("Ошибка запуска сервера: ", err)
	}
}
