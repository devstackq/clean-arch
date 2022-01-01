package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/devstackq/go-clean/auth/models"
	"github.com/devstackq/go-clean/auth/usecase"
	"github.com/stretchr/testify/assert"
)

//simple test 1 good case
func TestSignup(t *testing.T) {
	input := `{"username":"mock","password":"123s"}`

	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer([]byte(input)))
	if err != nil {
		t.Errorf("request error %v", err)
	}
	user := models.User{
		Username: "mock",
		Password: "123s",
	}

	mockService := usecase.AuthUseCaseMock{}

	mockService.Mock.On("SignUp", user.Username, user.Password).Return(nil)

	response := httptest.NewRecorder()
	hr := &Handler{}
	// hr.authUseCase = mockService
	handler := http.HandlerFunc(hr.SignUp)

	handler.ServeHTTP(response, req)

	assert.Equal(t, "", response.Body.String())
	assert.Equal(t, 200, response.Result().StatusCode)
}
