package repo

import "github.com/hiteshrepo/grpc-demo/internal/pkg/model"

type BookRepo interface {
	StoreBook(book *model.Book)
	GetBooks() []*model.Book
}

type bookRepo struct {
	booksList []*model.Book
}

func NewBookRepo() BookRepo {
	return &bookRepo{booksList: make([]*model.Book, 0)}
}

func (br *bookRepo) StoreBook(book *model.Book) {
	br.booksList = append(br.booksList, book)
}

func (br *bookRepo) GetBooks() []*model.Book {
	return br.booksList
}
