package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var password = "abcd12345"
var hashedPassword string
var errors error

func TestHashPassword(t *testing.T) {
	hashedPassword, errors = HashPassword(password)
	assert.NoError(t, errors, "hashing error")
	assert.NotEqual(t, password, hashedPassword, "password is not hashed")
}

func TestVerifyPassword(t *testing.T) {
	t.Run("verify success", func(t *testing.T) {
		var hashPassword = VerifyPassword(hashedPassword, password)
		assert.Nil(t, hashPassword, "wrong password")
	})

	t.Run("verify failed", func(t *testing.T) {
		var hashPassword = VerifyPassword(hashedPassword, "xxxxx")
		assert.NotNil(t, hashPassword, "password is correct")
	})
}
