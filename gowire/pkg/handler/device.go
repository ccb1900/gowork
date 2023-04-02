package handler

import (
	"context"

	"gowire/pkg/model"
	"gowire/pkg/repo"
)

type DeviceHandler struct {
	repo repo.IDeviceRepo
}

func (bh DeviceHandler) GetList(ctx context.Context) []model.Device {
	return bh.repo.Get()
}

func NewDeviceHandler(repo repo.IDeviceRepo) DeviceHandler {
	return DeviceHandler{
		repo: repo,
	}
}
