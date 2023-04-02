package repo

import (
	"gowire/pkg/model"
)

type IBookRepo interface {
	Get() []model.Book
}

type BookRepo struct{}

func (br BookRepo) Get() []model.Book {
	var books []model.Book

	for i := 0; i < 10; i++ {
		books = append(books, model.Book{
			Title: "hello",
		})
	}
	return books
}

func NewBookRepo() BookRepo {
	return BookRepo{}
}
