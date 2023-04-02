package handler

import (
	"context"

	"gowire/pkg/model"
	"gowire/pkg/repo"
)

type BookHandler struct {
	repo repo.IBookRepo
}

func (bh BookHandler) GetList(ctx context.Context) []model.Book {
	return bh.repo.Get()
}

func NewBookHandler(repo repo.IBookRepo) BookHandler {
	return BookHandler{
		repo: repo,
	}
}
