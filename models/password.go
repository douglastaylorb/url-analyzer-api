package models

type PasswordRequest struct {
	Length    int  `json:"length"`
	Uppercase bool `json:"uppercase"`
	Lowercase bool `json:"lowercase"`
	Numbers   bool `json:"numbers"`
	Special   bool `json:"special"`
}

type PasswordResponse struct {
	Password string `json:"password"`
}
