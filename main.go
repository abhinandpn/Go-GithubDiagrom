package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
)

const (
	formTemplatePath = "templates/form.html"
)

var (
	formDataMap = make(map[string]string)
	mu          sync.Mutex
)

func main() {
	http.HandleFunc("/", serveForm)
	http.HandleFunc("/submit", handleSubmit)
	http.HandleFunc("/data", getData)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(formTemplatePath))
	tmpl.Execute(w, nil)
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestBody map[string]string

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Log the data to the terminal
	log.Printf("Submitted Data: %+v\n", requestBody)

	// Send the personal data back to the client
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("Personal Data:\n%s", formatData(requestBody))))
}

func getData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	email := r.URL.Query().Get("email")
	mu.Lock()
	data, exists := formDataMap[email]
	mu.Unlock()

	if !exists {
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<pre>%s</pre>", data)
}

func formatData(data map[string]string) string {
	var result string
	for key, value := range data {
		result += fmt.Sprintf("%s: %s\n", key, value)
	}
	return result
}