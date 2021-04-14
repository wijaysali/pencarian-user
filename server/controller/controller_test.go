package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pencarian_user/server/service"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	getUsernameService func(urls []string) []string
)

type serviceMock struct{}

func (sm *serviceMock) UsernameCheck(urls []string) []string {
	return getUsernameService(urls)
}

func TestUsername_Success(t *testing.T) {
	service.UsernameService = &serviceMock{} //will now make our fake struct to implement the "usernameService" interface
	getUsernameService = func(urls []string) []string {
		return []string{
			"https://twitter.com/hacktiv8id",
			"https://dev.to/hacktiv8id",
			"https://instagram.com/hacktiv8id",
		}
	}
	r := gin.Default()
	jsonBody := `["https://twitter.com/hacktiv8id", "https://instagram.com/hacktiv8id", "https://dev.to/hacktiv8id"]`

	req, err := http.NewRequest(http.MethodPost, "/username", bytes.NewBufferString(jsonBody))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.POST("/username", Username)
	r.ServeHTTP(rr, req)

	var result []string
	err = json.Unmarshal(rr.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusOK, rr.Code)
	assert.EqualValues(t, 3, len(result))
}

//Here, we dont need to mock the service since we will never get there.
func TestUsername_Invalid_Data(t *testing.T) {
	r := gin.Default()
	//instead of using array syntax, we used object
	jsonBody := `{"https://twitter.com/hacktiv8id", "https://instagram.com/hacktiv8id", "https://github.com/hacktiv8id"}`

	req, err := http.NewRequest(http.MethodPost, "/username", bytes.NewBufferString(jsonBody))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.POST("/username", Username)
	r.ServeHTTP(rr, req)

	var result []string
	err = json.Unmarshal(rr.Body.Bytes(), &result)

	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.EqualValues(t, http.StatusUnprocessableEntity, rr.Code)
}
