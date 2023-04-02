package handler

import (
	"context"
	"log"

	"gowire/pkg/model"
	"gowire/pkg/repo"
)

type BookHandler struct {
	book   repo.IBookRepo
	device repo.IDeviceRepo
}

func (bh BookHandler) GetList(ctx context.Context) []model.Book {
	log.Println("设备", bh.device.Get())
	return bh.book.Get()
}

func NewBookHandler(book repo.IBookRepo, device repo.IDeviceRepo) BookHandler {
	return BookHandler{
		book:   book,
		device: device,
	}
}
