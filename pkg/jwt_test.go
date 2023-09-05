package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var uid string = "user123"
var role string = "admin"

func TestNewToken(t *testing.T) {
	token := NewToken(uid, role)

	// Add assertions to verify the content of the token
	assert.Equal(t, uid, token.Id, "Expected UID to match")
	assert.Equal(t, role, token.Role, "Expected role to match")
	assert.NotEmpty(t, token.ExpiresAt, "Expected ExpiresAt to be set")
	assert.Equal(t, "coffeshop_backend", token.Issuer, "Expected Issuer to match")
}

func TestGenerateToken(t *testing.T) {
	token := NewToken(uid, role)

	tokenString, err := token.Generate()

	assert.NoError(t, err, "Expected no error while generating token")
	assert.NotEmpty(t, tokenString, "Expected non-empty token string")
}

func TestVerifyToken(t *testing.T) {
	token := NewToken(uid, role)
	tokenString, _ := token.Generate()

	verifiedToken, err := VerifyToken(tokenString)

	assert.NoError(t, err, "Expected no error while verifying token")
	assert.Equal(t, uid, verifiedToken.Id, "Expected UID to match")
	assert.Equal(t, role, verifiedToken.Role, "Expected role to match")
	assert.Equal(t, "coffeshop_backend", verifiedToken.Issuer, "Expected Issuer to match")
}
