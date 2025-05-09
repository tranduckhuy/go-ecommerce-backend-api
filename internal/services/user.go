package services

import "github.com/tranduckhuy/go-ecommerce-backend-api/internal/repositories"

// type UserService struct {
// 	userRepository *repositories.UserRepository
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepository: repositories.NewUserRepository(),
// 	}
// }

// func (us *UserService) GetUserByID(userID string) (string, error) {
// 	// Simulate a user retrieval
// 	return us.userRepository.GetUserByID(userID)
// }

type UserService interface {
	GetUserByID(userID string) (string, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (us *userService) GetUserByID(userID string) (string, error) {
	// Simulate a user retrieval
	return us.userRepository.GetUserByID(userID)
}
