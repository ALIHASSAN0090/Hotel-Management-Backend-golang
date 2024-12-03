package database_repo

type AiRepository interface {
	GetAiResponceDB(order_details, hotel_details, question string) (string, error)
}
