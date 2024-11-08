package tests

import (
	"Eight/src/internal/customError"
	"Eight/src/internal/domain"
	"Eight/src/payload/response"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	myHttp "Eight/src/internal/delivery/http"
	"github.com/labstack/echo/v4"
)

// MockUserService - это мок-сервис для тестирования
type MockUserService struct {
	users map[int]*domain.User
}

func (m *MockUserService) Create(user *domain.User) error {
	m.users[int(user.ID)] = user
	return nil
}

func (m *MockUserService) FindAllWithPagination(offset, limit int, name string, age int) ([]domain.User, error) {
	var users []domain.User
	for _, user := range m.users {
		users = append(users, *user)
	}
	return users, nil
}

func (m *MockUserService) FindById(id int) (domain.User, error) {
	user, exists := m.users[id]
	if !exists {
		return domain.User{}, customError.New(http.StatusNotFound, "User not found")
	}
	return *user, nil
}

func (m *MockUserService) Update(user *domain.User) error {
	m.users[int(user.ID)] = user
	return nil
}

func (m *MockUserService) Delete(id uint) error {
	delete(m.users, int(id))
	return nil
}

func TestUserHandler(t *testing.T) {
	e := echo.New()
	mockService := &MockUserService{users: make(map[int]*domain.User)}
	userHandler := myHttp.NewUserHandler(mockService)

	e.POST("/api/users", userHandler.Create)
	e.GET("/api/users", userHandler.GetAll)
	e.GET("/api/users", userHandler.GetById)
	e.PUT("/api/users", userHandler.Update)
	e.DELETE("/api/users", userHandler.Delete)

	t.Run("Create User", func(t *testing.T) {
		user := &domain.User{
			ID:       1,
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password",
			Birthday: time.Now(),
		}
		body, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if err := userHandler.Create(c); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if rec.Code != http.StatusCreated {
			t.Errorf("Expected status 201, got %v", rec.Code)
		}
	})

	t.Run("Get All Users", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if err := userHandler.GetAll(c); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %v", rec.Code)
		}
	})

	t.Run("Get User By ID", func(t *testing.T) {
		user := &domain.User{
			ID:       1,
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password",
			Birthday: time.Now(),
		}
		mockService.Create(user)

		req := httptest.NewRequest(http.MethodGet, "/api/users?id=1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if err := userHandler.GetById(c); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %v", rec.Code)
		}

		var responseUser domain.User
		if err := json.NewDecoder(rec.Body).Decode(&responseUser); err != nil {
			t.Errorf("Failed to decode response: %v", err)
		}

		if responseUser.ID != user.ID {
			t.Errorf("Expected user ID %d, got %d", user.ID, responseUser.ID)
		}
	})

	t.Run("Update User", func(t *testing.T) {
		user := &domain.User{
			ID:       1,
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password",
			Birthday: time.Now(),
		}
		mockService.Create(user)

		updates := &response.UpdateResponse{
			Name:     "Jane Doe",
			Email:    "jane@example.com",
			Password: "newpassword",
			Birthday: time.Now(),
		}
		body, _ := json.Marshal(updates)
		req := httptest.NewRequest(http.MethodPut, "/api/users?id=1", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if err := userHandler.Update(c); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if rec.Code != http.StatusNoContent {
			t.Errorf("Expected status 204, got %v", rec.Code)
		}

		// Проверяем, что пользователь обновился
		updatedUser, _ := mockService.FindById(1)
		if updatedUser.Name != updates.Name {
			t.Errorf("Expected user name %s, got %s", updates.Name, updatedUser.Name)
		}
	})

	t.Run("Delete User", func(t *testing.T) {
		user := &domain.User{
			ID:       1,
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password",
			Birthday: time.Now(),
		}
		mockService.Create(user)

		req := httptest.NewRequest(http.MethodDelete, "/api/users?id=1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if err := userHandler.Delete(c); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if rec.Code != http.StatusNoContent {
			t.Errorf("Expected status 204, got %v", rec.Code)
		}

		_, err := mockService.FindById(1)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
