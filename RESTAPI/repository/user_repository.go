// repository/user.go
package repository

import (
	"restapi.com/myuser/myapi/mockdata"
	"restapi.com/myuser/myapi/models"
)

// UserRepository interface
type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int64) (models.User, error)
	CreateUser(user models.User) (int64, error)
	UpdateUser(id int64, user models.User) error
	DeleteUser(id int64) error
}

// Mock data to use during testing
// UserRepository struct
type userRepository struct {
}

// NewUserRepository function returns a new UserRepository instance with mock data
func NewUserRepository() UserRepository {
	return &userRepository{}
}

// GetAllUsers function returns all users
func (*userRepository) GetAllUsers() ([]models.User, error) {
	return mockdata.Users, nil
}

// GetUserByID function returns a user by ID
func (*userRepository) GetUserByID(id int64) (models.User, error) {
	for _, user := range mockdata.Users {
		if user.ID == id {
			return user, nil
		}
	}

	return models.User{}, nil
}

// CreateUser function creates a new user
func (*userRepository) CreateUser(user models.User) (int64, error) {
	user.ID = int64(len(mockdata.Users) + 1)
	mockdata.Users = append(mockdata.Users, user)

	return user.ID, nil
}

// UpdateUser function updates an existing user
func (*userRepository) UpdateUser(id int64, user models.User) error {
	for i, u := range mockdata.Users {
		if u.ID == id {
			user.ID = id
			mockdata.Users[i] = user

			return nil
		}
	}

	return nil
}

// DeleteUser function deletes a user by ID
func (r *userRepository) DeleteUser(id int64) error {
	for i, user := range mockdata.Users {
		if user.ID == id {
			mockdata.Users = append(mockdata.Users[:i], mockdata.Users[i+1:]...)
			break
		}
	}

	return nil
}
