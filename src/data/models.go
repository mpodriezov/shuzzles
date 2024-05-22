package data

import "time"

type RegistrationModel struct {
	Username     string
	Email        string
	PasswordHash string
	Bio          string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserModel struct {
	Id        int
	Username  string
	Email     string
	Bio       string
	Altitude  float64
	Longitude float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
