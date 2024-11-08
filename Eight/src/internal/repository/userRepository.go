package repository

import (
	"Eight/src/internal/domain"
	"gorm.io/gorm"
	"time"
)

type UserRepositoryInterface interface {
	Create(user *domain.User) error
	FindAllWithPagination(offset, limit int, name string, age int) ([]domain.User, error)
	FindById(id int) (domain.User, error)
	Update(user *domain.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{db: db}
}

func (r *userRepository) Create(todo *domain.User) error {

	return r.db.Create(todo).Error
}

func (r *userRepository) FindAllWithPagination(offset, limit int, name string, age int) ([]domain.User, error) {
	var users []domain.User
	query := r.db.Offset(offset).Limit(limit)

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if age > 0 {
		birthday := time.Now().AddDate(-age, 0, 0)
		query = query.Where("birthday <= ?", birthday)
	}

	err := query.Find(&users).Error
	return users, err
}

func (r *userRepository) FindById(id int) (domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *userRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}
