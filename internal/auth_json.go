package internal

import (
	"encoding/json"
	"time"
)

// AuthJson represents a request sent to /users/register and /users/login
type AuthJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewPasswordJson represents a request sent to /users/new-password
type NewPasswordJson struct {
	NewPassword string `json:"newPassword"`
}

type LoginSuccessful struct {
	Status     string `json:"status"`
	Username   string `json:"username"`
	UserUuid   string `json:"userUuid"`
	Expiration string `json:"expiration"`
	Token      string `json:"token"`
}

func (s *LoginSuccessful) Marshal() []byte {
	j, _ := json.Marshal(s)
	return j
}

func LoginResponse(
	resp struct {
		Status   string
		Username string
		UserUuid string
		Exp      time.Time
		Token    string
	},
) *LoginSuccessful {
	return &LoginSuccessful{
		Status:     resp.Status,
		Username:   resp.Username,
		UserUuid:   resp.UserUuid,
		Expiration: resp.Exp.String(),
		Token:      resp.Token,
	}
}
