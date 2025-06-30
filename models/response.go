package models

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message,omitempty"`
	Results any `json:"results,omitempty"`
	Errors	any `json:"errors,omitempty"`
}

