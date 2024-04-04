package models

type AIService struct {
	UID         string
	Title       string
	Description string
	Price       float64
}

type AIServiceCreate struct {
	Title       string
	Description *string
	Price       float64
}
