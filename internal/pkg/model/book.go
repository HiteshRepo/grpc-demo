package model

type Book struct {
	ISBN   int    `json:"isbn" db:"isbn"`
	Name   string `json:"name" db:"name"`
	Author string `json:"author" db:"author"`
}

func GetNewBook(isbn int, name, author string) *Book {
	return &Book{ISBN: isbn, Name: name, Author: author}
}
