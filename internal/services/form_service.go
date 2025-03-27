package services

import (
	"sync"

	models "github.com/abhinandpn/Go-Html-Phishing/internal/model"
)

var (
	formDataStore []models.FormData
	mu            sync.Mutex
)

func SaveFormData(data models.FormData) {
	mu.Lock()
	defer mu.Unlock()
	formDataStore = append(formDataStore, data)
}

func GetAllFormData() []models.FormData {
	mu.Lock()
	defer mu.Unlock()
	return formDataStore
}
