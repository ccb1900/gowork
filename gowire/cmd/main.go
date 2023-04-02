package main

import (
	"context"
	"log"

	"gowire/pkg/di"
)

func main() {
	book := di.InitializeBookHandler()
	device := di.InitializeDeviceHandler()
	log.Println(device.GetList(context.Background()))
	log.Println(book.GetList(context.Background()))
}
