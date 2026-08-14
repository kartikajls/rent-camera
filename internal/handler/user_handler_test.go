package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Login(email, password string) (string, error) {
	args := m.Called(email, password)

	return args.String(0), args.Error(1)
}

func TestLogin(t *testing.T) {
	mockUsecase := new(MockUserUsecase)

	// Data yang diharapkan
	mockUsecase.
		On("Login", "sasuke@example.com", "123456").
		Return("dummy-token", nil)

	// Data yang digunakan untuk test
	// ini yang akan di adjust ketika di test
	email := "sasuke1@example.com"
	password := "123456"

	// Jalankan mock
	token, err := mockUsecase.Login(email, password)

	// Assertion
	assert.NoError(t, err)
	assert.Equal(t, "dummy-token", token)

	// Pastikan mock sesuai dengan expectation
	mockUsecase.AssertExpectations(t)
}
