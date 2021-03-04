package main

type advertisement struct {
	ID          int
	Active      bool
	CreatedOn   string
	UserID      int
	Name        string
	Category    string
	Price       string
	Description string
	Images      []image
}

type image struct {
	ID      int
	Path    string
	ProdRef int
}
