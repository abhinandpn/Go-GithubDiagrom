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

	r.ParseForm()
	resultText := "You submitted the following data:<br>"
	for key, values := range r.Form {
		for _, value := range values {
			resultText += fmt.Sprintf("<strong>%s:</strong> %s<br>", key, value)
		}
	}

	ipAddress, err := getClientIP()
	if err != nil {
		ipAddress = "Unable to retrieve IP"
	}
	resultText += fmt.Sprintf("<strong>IP Address:</strong> %s<br>", ipAddress)

	// Save the data to the in-memory map
	mu.Lock()
	formDataMap[r.FormValue("email")] = resultText
	mu.Unlock()

	// Log the data to the terminal
	log.Printf("Submitted Data: %s\n", resultText)

	tmpl := template.Must(template.ParseFiles(formTemplatePath))
	tmpl.Execute(w, struct {
		Result string
	}{
		Result: resultText,
	})
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

func getClientIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		IP string `json:"ip"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.IP, nil
}