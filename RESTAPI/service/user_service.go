// models/user.go
package service

import (
	"restapi.com/myuser/myapi/models"
	"restapi.com/myuser/myapi/repository"
)

// UserRepository interface
type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int64) (models.User, error)
	CreateUser(user models.User) (int64, error)
	UpdateUser(id int64, user models.User) error
	DeleteUser(id int64) error
}

// UserService struct
type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

var (
	s repository.UserRepository = repository.NewUserRepository()
)

// GetUsers function returns all users
func (*userService) GetAllUsers() ([]models.User, error) {
	return s.GetAllUsers()
}

// GetUserByID function returns a user by ID
func (*userService) GetUserByID(id int64) (models.User, error) {
	return s.GetUserByID(id)
}

// CreateUser function creates a new user
func (*userService) CreateUser(user models.User) (int64, error) {
	return s.CreateUser(user)
}

// UpdateUser function updates an existing user
func (*userService) UpdateUser(id int64, user models.User) error {
	return s.UpdateUser(id, user)
}

// DeleteUser function deletes a user by ID
func (*userService) DeleteUser(id int64) error {
	return s.DeleteUser(id)
}
