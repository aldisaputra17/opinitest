package service

import (
	"log"

	"github.com/aldisaputra17/post_opinia/dataobject"
	"github.com/aldisaputra17/post_opinia/models"
	"github.com/aldisaputra17/post_opinia/repository"
	"github.com/mashingan/smapping"
)

type UserService interface {
	Update(user dataobject.UserUpdateObj) models.User
	Profile(userID string) models.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dataobject.UserUpdateObj) models.User {
	userToUpdate := models.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) models.User {
	return service.userRepository.ProfileUser(userID)
}
