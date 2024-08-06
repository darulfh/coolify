package dto

type PulsaDto struct {
	Type        string `json:"type"`
	Provider    string `json:"provider"`
	PhoneNumber string `json:"phone_number"`
	Limit       int    `json:"limit"`
	Page        int    `json:"page"`
}
