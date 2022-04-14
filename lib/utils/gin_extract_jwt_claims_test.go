package utils

import (
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/artnoi43/todong/test"
)

var (
	testContext = &gin.Context{
		Keys: test.JwtKeys,
	}
)

func TestExtractJwtClaims(t *testing.T) {
	userInfo, err := ExtractAndDecodeJwt(testContext)
	if err != nil {
		t.Fatal(err)
	}
	if userInfo.UserUuid != test.JwtIss {
		t.Logf("expected: %s\n", test.JwtIss)
		t.Logf("actual: %s\n", userInfo.UserUuid)
		t.Fatal("invalid iss")
	}
	if userInfo.Expiration != test.JwtExp.Local() {
		t.Logf("expected: %s\n", test.JwtExp)
		t.Logf("actual: %s\n", userInfo.Expiration)
		t.Fatal("invalid exp")
	}
}
