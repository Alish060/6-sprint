package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "server: ", log.LstdFlags)

	s := server.CreateRouter(logger)

	logger.Println("Запуск сервера")
	err := s.Server.ListenAndServe()
	if err != nil {
		logger.Fatal("Ошибка при запуске сервера:", err)
	}
}
