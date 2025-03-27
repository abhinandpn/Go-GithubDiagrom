package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abhinandpn/Go-Html-Phishing/config"
	handlers "github.com/abhinandpn/Go-Html-Phishing/internal/handler"
)

func main() {
	config.LoadConfig()

	http.HandleFunc("/submit-form", handlers.GetAllForms)
	http.Handle("/", http.FileServer(http.Dir("static"))) // Serves HTML form

	port := ":" + config.AppConfig.ServerPort
	fmt.Println("Server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
