package domain

type GetUserRequest struct {
	PassportSerie  string `json:"passport_serie"`
	PassportNumber string `json:"passport_number"`
}

type GetUserResponse struct {
	PassportSerie  string `json:"passport_credentials"`
	PassportNumber string `json:"passport_number"`
	People         People
}

type People struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
	Address    string `json:"address"`
}

type UsersListParams struct {
	PassportSerie  string `form:"passport_serie"`
	PassportNumber string `form:"passport_number"`
	Name           string `form:"name"`
	Surname        string `form:"surname"`
	Patronymic     string `form:"patronymic"`
	Address        string `form:"address"`
	Limit          int    `form:"limit"`
	Offset         int    `form:"offset"`
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
