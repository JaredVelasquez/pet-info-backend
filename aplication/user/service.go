package user

import (
	model "pet-info-backend-docker/domain/user"
	repository "pet-info-backend-docker/infraestructure/repository/user"
)

// UserService define los métodos de negocio para los usuarios
type UserService struct {
	Repo *repository.UserRepository
}

// NewUserService crea una nueva instancia de UserService
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// CreateUser crea un nuevo usuario
func (service *UserService) CreateUser(user *model.User) error {
	return service.Repo.CreateUser(user)
}

// GetUser obtiene un usuario por email o username
func (service *UserService) GetUser(email, username string) ([]model.User, error) {
	if email != "" {
		return service.Repo.GetUserByEmail(email)
	} else if username != "" {
		return service.Repo.GetUserByUsername(username)
	}
	return nil, nil
}
