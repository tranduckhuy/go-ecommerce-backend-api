package services

import "github.com/tranduckhuy/go-ecommerce-backend-api/internal/repositories"

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repositories.NewUserRepository(),
	}
}

func (us *UserService) GetUserByID(userID string) (string, error) {
	// Simulate a user retrieval
	return us.userRepository.GetUserByID(userID)
}
