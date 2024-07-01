package service

import (
	"github.com/zsandibe/eff_mobile_task/config"
	"github.com/zsandibe/eff_mobile_task/internal/repository"
)

type Service interface {
}

type serviceCar struct {
	repo repository.Repository
	conf config.Config
}

func NewService(repo repository.Repository) *serviceCar {
	return &serviceCar{
		repo: repo,
	}
}
