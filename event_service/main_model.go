package main

type event struct {
	ID          int
	Active      bool
	CreatedOn   string
	UserID      int
	ProductID   int `json:"product_id"`
	Name        string
	Category    string
	From        string
	To          string
	Place       string
	Description string
	Images      []image
}

type image struct {
	ID       int
	Path     string
	EventRef int
}
