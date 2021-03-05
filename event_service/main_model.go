package main

type event struct {
	ID          int     `json:"ID"`
	Active      bool    `json:"Active"`
	CreatedOn   string  `json:"CreatedOn"`
	UserID      int     `json:"UserID"`
	ProductID   int     `json:"ProductID"`
	Name        string  `json:"Name"`
	Category    string  `json:"Category"`
	From        string  `json:"From"`
	To          string  `json:"To"`
	Place       string  `json:"Place"`
	Description string  `json:"Description"`
	Images      []image `json:"Images"`
}

type image struct {
	ID       int
	Path     string
	EventRef int
}
