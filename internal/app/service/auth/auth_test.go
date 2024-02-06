package auth

import (
	"github.com/chenxingyuu/gin_template/internal/app/model"
	"testing"
)

func TestAuthenticateWithIncorrectPassword(t *testing.T) {
	user := &model.User{Password: "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"} // bcrypt hash of "password"
	if Authenticate(user, "wrong password") {
		t.Errorf("Authenticate with incorrect password passed")
	}
}

func TestEncryptPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := EncryptPassword(password)
	if err != nil {
		t.Errorf("EncryptPassword failed with error: %v", err)
	}

	user := &model.User{Password: hashedPassword}
	if !Authenticate(user, password) {
		t.Errorf("EncryptPassword did not produce a valid bcrypt hash")
	}
}
