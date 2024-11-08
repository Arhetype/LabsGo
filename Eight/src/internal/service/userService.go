package service

import (
	"Eight/src/internal/domain"
	"Eight/src/internal/repository"
	"github.com/go-playground/validator/v10"
)

type UserServiceInterface interface {
	Create(user *domain.User) error
	FindAllWithPagination(offset, limit int, name string, age int) ([]domain.User, error)
	FindById(id int) (domain.User, error)
	Update(user *domain.User) error
	Delete(id uint) error
}

type UserService struct {
	repo      repository.UserRepositoryInterface
	validator *validator.Validate
}

func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{
		repo:      repo,
		validator: validator.New(),
	}
}

func (service *UserService) Create(user *domain.User) error {
	if err := service.validator.Struct(user); err != nil {
		return err
	}
	return service.repo.Create(user)
}

func (service *UserService) FindAllWithPagination(offset, limit int, name string, age int) ([]domain.User, error) {
	return service.repo.FindAllWithPagination(offset, limit, name, age)
}

func (service *UserService) FindById(id int) (domain.User, error) {
	return service.repo.FindById(id)
}

func (service *UserService) Update(user *domain.User) error {
	return service.repo.Update(user)
}

func (service *UserService) Delete(id uint) error {
	return service.repo.Delete(id)
}
