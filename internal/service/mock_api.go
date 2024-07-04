package service

import (
	"fmt"

	"github.com/zsandibe/eff_mobile_task/internal/domain"
)

func getInfoByPassport(passportSerie, passportNumber string) (domain.GetUserResponse, error) {
	fmt.Println(passportSerie, passportNumber)
	mockResponse := domain.GetUserResponse{
		PassportSerie:  passportSerie,
		PassportNumber: passportNumber,
		People: domain.People{
			Name:       "Тест",
			Surname:    "Тестов",
			Patronymic: "Тестович",
			Address:    "Тестовенко 14",
		},
	}
	return mockResponse, nil
}
