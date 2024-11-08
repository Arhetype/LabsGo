package http

import (
	"Eight/src/internal/customError"
	"Eight/src/internal/domain"
	"Eight/src/internal/service"
	"Eight/src/payload/response"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserHandlerInterface interface {
	Create(c echo.Context) error
	GetAll(c echo.Context) error
	GetByID(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type UserHandler struct {
	service service.UserServiceInterface
}

func NewUserHandler(service service.UserServiceInterface) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Create(c echo.Context) error {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return customError.New(http.StatusBadRequest, "Invalid input")
	}

	if err := h.service.Create(&user); err != nil {
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) {
			return customError.New(http.StatusBadRequest, "Validation failed: "+validationErrs.Error())
		}
		return customError.New(http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetAll(c echo.Context) error {
	idParam := c.QueryParam("id")
	if idParam != "" {
		return h.GetById(c)
	}
	offsetParam := c.QueryParam("offset")
	limitParam := c.QueryParam("limit")
	nameParam := c.QueryParam("name")
	ageParam := c.QueryParam("age")

	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		limit = 10
	}

	var age int
	if ageParam != "" {
		age, err = strconv.Atoi(ageParam)
		if err != nil {
			return customError.New(http.StatusBadRequest, "Invalid age parameter")
		}
	}

	users, err := h.service.FindAllWithPagination(offset, limit, nameParam, age)
	if err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	var userResponse []response.UserResponse
	total := len(users)

	for _, user := range users {
		userResponse = append(userResponse, response.NewUserResponse(user))
	}

	usersResponse := response.UsersResponse{
		Total:  total,
		Offset: offset,
		Limit:  limit,
		Users:  userResponse,
	}

	return c.JSON(http.StatusOK, usersResponse)
}

func (h *UserHandler) GetById(c echo.Context) error {
	idParam := c.QueryParam("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return customError.New(http.StatusBadRequest, "Invalid id parameter")
	}
	user, err := h.service.FindById(id)
	if err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}
	userResponse := response.NewUserResponse(user)
	return c.JSON(http.StatusOK, userResponse)
}

func (h *UserHandler) Update(c echo.Context) error {
	idStr := c.QueryParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return customError.New(http.StatusBadRequest, "Invalid id parameter")
	}

	var updates response.UpdateResponse

	if err := c.Bind(&updates); err != nil {
		return customError.New(http.StatusBadRequest, err.Error())
	}

	user, err := h.service.FindById(id)
	if err != nil {
		return customError.New(http.StatusBadRequest, err.Error())
	}

	user.Name = updates.Name
	user.Email = updates.Email
	user.Password = updates.Password
	user.Birthday = updates.Birthday

	if err := h.service.Update(&user); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (h *UserHandler) Delete(c echo.Context) error {
	id, _ := strconv.ParseUint(c.QueryParam("id"), 10, 32)
	if err := h.service.Delete(uint(id)); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
