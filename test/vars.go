package test

import (
	"encoding/json"
	"time"

	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/artnoi43/todong/internal"
	"golang.org/x/crypto/bcrypt"
)

var (
	Username    = "torvalds"
	Password    = []byte("stallman")
	HashedPW, _ = bcrypt.GenerateFromPassword(Password, 14)
	JwtIss      = "0x69"
	JwtExp, _   = time.Parse(time.RFC3339, time.RFC3339)
	JwtKeys     = map[string]interface{}{
		"JWT_PAYLOAD": ginjwt.MapClaims{
			"iss": JwtIss,
			"exp": float64(JwtExp.Local().Unix()),
		},
	}

	TodoUuid        = "cdb1eb67-fe8c-44ba-a32a-bda8ca95bd3c"
	TodoTitle       = "Testing title"
	TodoDescription = "Write tests first!"
	TodoDate        = "2021-01-01"
	Todo            = internal.TodoReqBody{
		Title:       TodoTitle,
		Description: TodoDescription,
		TodoDate:    TodoDate,
	}
	TodoJson, _ = json.Marshal(Todo)
)
