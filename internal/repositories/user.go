package repositories

import "fmt"

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) GetUserByID(userID string) (string, error) {
	// Simulate a user retrieval
	if userID == "1" {
		return "User 1", nil
	}
	return "", fmt.Errorf("user not found")
}
