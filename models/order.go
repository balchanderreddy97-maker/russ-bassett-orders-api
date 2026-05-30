package models

type Order struct {
	ID           int    `json:"id"`
	ClientName   string `json:"client_name"`
	ProjectType  string `json:"project_type"`
	Status       string `json:"status"`
	DeliveryDate string `json:"delivery_date"`
}