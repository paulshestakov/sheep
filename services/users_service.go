package services

import (
	"main/domain"
	"main/repositories"
)


type usersService struct{
	repository repositories.UsersRepositoryInterface
}

type UsersServiceInterface interface {
	Get(string) (*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
}

func NewUsersService(repository repositories.UsersRepositoryInterface) UsersServiceInterface {
	return &usersService{repository: repository}
}

func (u *usersService) Get(id string) (*domain.User, error) {
	return u.repository.Get(id)
}

func (u *usersService) Create(user *domain.User) (*domain.User, error) {
	return u.repository.Create(user)
}