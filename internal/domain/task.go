package domain

type CreateTaskRequest struct {
	UserId      string `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type StopTaskRequest struct {
	UserId string `json:"user_id"`
}
