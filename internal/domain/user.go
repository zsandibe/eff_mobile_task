package domain

type GetUserRequest struct {
	PassportSerie  string `json:"passport_serie"`
	PassportNumber string `json:"passport_number"`
}

type GetUserResponse struct {
	PassportCredentials GetUserRequest `json:"passport_credentials"`
	Name                string         `json:"name"`
	Surname             string         `json:"surname"`
	Patronymic          string         `json:"patronymic,omitempty"`
	Address             string         `json:"address"`
}

type UsersListParams struct {
	Id                  int            `json:"id"`
	PassportCredentials GetUserRequest `json:"passport_credentials"`
	Name                string         `json:"name"`
	Surname             string         `json:"surname"`
	Patronymic          string         `json:"patronymic"`
	Address             string         `json:"address"`
	Limit               int            `json:"limit"`
	Offset              int            `json:"offset"`
}

type UserDataUpdatingRequest struct {
	Id             int    `json:"id"`
	PassportSerie  string `json:"passport_serie"`
	PassportNumber string `json:"passport_number"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
}
