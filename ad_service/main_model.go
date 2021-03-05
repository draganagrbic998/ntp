package main

type advertisement struct {
	ID          int     `json:"ID"`
	Active      bool    `json:"Active"`
	CreatedOn   string  `json:"CreatedOn"`
	UserID      int     `json:"UserID"`
	Name        string  `json:"Name"`
	Category    string  `json:"Category"`
	Price       int     `json:"Price"`
	Description string  `json:"Description"`
	Images      []image `json:"Images"`
}

type image struct {
	ID      int
	Path    string
	ProdRef int
}
