package v1

import "github.com/zsandibe/eff_mobile_task/internal/service"

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{service: service}
}
