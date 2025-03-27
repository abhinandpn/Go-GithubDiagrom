package models

type FormData struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Address         string `json:"address"`
	City            string `json:"city"`
	State           string `json:"state"`
	Country         string `json:"country"`
	ZipCode         string `json:"zip_code"`
	Company         string `json:"company"`
	Position        string `json:"position"`
	Message         string `json:"message"`
	AdditionalNotes string `json:"additional_notes"`
}
