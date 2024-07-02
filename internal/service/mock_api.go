package service

import (
	"fmt"

	"github.com/zsandibe/eff_mobile_task/internal/domain"
)

func getInfoByPassport(passportSerie, passportNumber string) (domain.GetUserResponse, error) {
	fmt.Println(passportSerie, passportNumber)
	mockResponse := domain.GetUserResponse{
		PassportCredentials: domain.GetUserRequest{
			PassportSerie:  passportSerie,
			PassportNumber: passportNumber,
		},
		Name:       "DAUN",
		Surname:    "Малов",
		Patronymic: "GANDON",
		Address:    "Атаманонова 14",
	}
	return mockResponse, nil
}
