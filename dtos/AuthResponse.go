package dtos

import "time"

type AuthResponse struct {
	Token string    `json:"token"`
	Exp   time.Time `json:"exp"`
}
