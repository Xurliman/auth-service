package services

import (
	"github.com/Xurliman/auth-service/internal/server/app/dto"
	"github.com/Xurliman/auth-service/internal/server/app/interfaces"
)

type AuthService struct {
	repository interfaces.IAuthRepository
}

func (s *AuthService) Register(request *dto.RegisterRequest) error {
	return s.repository.Insert(request.ToModel())
}

func (s *AuthService) Login(request *dto.LoginRequest) error {
	return s.repository.Insert(request.ToModel())
}

func NewAuthService(repository interfaces.IAuthRepository) interfaces.IAuthService {
	return &AuthService{
		repository: repository,
	}
}
