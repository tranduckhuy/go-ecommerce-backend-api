package repositories

import "fmt"

// type UserRepository struct{}

// func NewUserRepository() *UserRepository {
// 	return &UserRepository{}
// }

// func (ur *UserRepository) GetUserByID(userID string) (string, error) {
// 	// Simulate a user retrieval
// 	if userID == "1" {
// 		return "User 1", nil
// 	}
// 	return "", fmt.Errorf("user not found")
// }

type UserRepository interface {
	GetUserByID(userID string) (string, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (ur *userRepository) GetUserByID(userID string) (string, error) {
	// Simulate a user retrieval
	if userID == "1" {
		return "User 1-1", nil
	}
	return "", fmt.Errorf("user not found")
}
