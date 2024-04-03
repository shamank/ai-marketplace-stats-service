package models

type AIService struct {
	UID         string
	Title       string
	Description string
	Price       int
}

type AIServiceCreate struct {
	Title       string
	Description *string
	Price       int
}
