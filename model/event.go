package model

const (
	Sold 	= "sold"
	Bought	= "bought"
)

type Event struct {
	Name 	string		`json:"event_name"`
	Product *Product	`json:"product"`
}