package repo

import (
	"gowire/pkg/model"
)

type IDeviceRepo interface {
	Get() []model.Device
}

type DeviceRepo struct{}

func (br DeviceRepo) Get() []model.Device {
	var books []model.Device

	for i := 0; i < 10; i++ {
		books = append(books, model.Device{
			Name: "world",
		})
	}
	return books
}

func NewDeviceRepo() DeviceRepo {
	return DeviceRepo{}
}
