package entity

type User struct {
	Id             int    `json:"id" db:"id"`
	PassportSerie  string `json:"passport_serie" db:"passport_serie"`
	PassportNumber string `json:"passport_number" db:"passport_number"`
	Name           string `json:"name" db:"name"`
	Surname        string `json:"surname" db:"surname"`
	Patronymic     string `json:"patronymic,omitempty" db:"patronymic"`
	Address        string `json:"address" db:"address"`
}
