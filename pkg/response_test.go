package pkg

import (
	"fmt"
	"gogin/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRes(t *testing.T) {
	code := 200
	data := &config.Result{}

	response := NewRes(code, data)

	assert.Equal(t, code, response.Code, "Expected response code to match")
	assert.Equal(t, getStatus(code), response.Status, "Expected response status to match")
	assert.Equal(t, data.Data, response.Data, "Expected response data to match")
	assert.Nil(t, response.Meta, "Expected response meta to be nil")
	assert.Nil(t, response.Description, "Expected response description to be nil")
}

func TestGetStatus(t *testing.T) {
	testCases := []struct {
		code   int
		status string
	}{
		{200, "OK"},
		{201, "Created"},
		{400, "Bad Request"},
		{401, "Unauthorized"},
		{403, "Forbidden"},
		{404, "Not Found"},
		{500, "Internal Server Error"},
		{501, "Bad Gateway"},
		{304, "Not Modified"},
		// Add more test cases as needed.
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("getStatus(%d)", tc.code), func(t *testing.T) {
			status := getStatus(tc.code)
			assert.Equal(t, tc.status, status, "Expected status to match")
		})
	}
}
