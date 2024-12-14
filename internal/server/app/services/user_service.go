package services

import (
	"errors"
	"github.com/Xurliman/auth-service/internal/config/config"
	"github.com/Xurliman/auth-service/internal/constants"
	"github.com/Xurliman/auth-service/internal/server/app/interfaces"
	"github.com/Xurliman/auth-service/internal/server/app/mail"
	"github.com/Xurliman/auth-service/internal/server/app/middlewares"
	"github.com/Xurliman/auth-service/internal/server/app/requests"
	"github.com/Xurliman/auth-service/internal/server/app/responses"
	"github.com/Xurliman/auth-service/pkg/json"
	"github.com/Xurliman/auth-service/pkg/pagination"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
	repository interfaces.IUserRepository
}

func NewUserService(repository interfaces.IUserRepository) interfaces.IUserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) Update(ctx *fiber.Ctx, id string, request *requests.UpdateUserRequest) error {
	if err := request.Validate(); err != nil {
		return json.ErrorValidation(ctx, err)
	}

	_, err := s.repository.FindById(id)
	if err != nil {
		return json.ErrorNotFound(ctx, err)
	}

	_, err = s.repository.UpdateById(id, request.ToModel())
	if err != nil {
		return json.Error(ctx, err, "ERR_UPDATE_USER")
	}

	return json.Success(ctx, fiber.Map{"id": id})
}

func (s *UserService) Add(ctx *fiber.Ctx, request *requests.StoreUserRequest) error {
	if err := request.Validate(); err != nil {
		return json.ErrorValidation(ctx, err)
	}

	if s.repository.EmailExists(request.Email) {
		return json.ErrorValidation(ctx, constants.ErrEmailExists)
	}

	token, err := s.generateEmailVerificationToken(request.Email)
	if err != nil {
		return json.Error(ctx, err, "ERR_GENERATING_EMAIL_VERIFICATION_TOKEN")
	}

	verificationUrl := ctx.BaseURL() + "/api/auth/verify-email?token=" + token
	go mail.VerifyEmail(request.Email, verificationUrl)

	user, err := s.repository.Create(request.ToModel())
	if err != nil {
		return json.Error(ctx, err, "ERR_CREATE_USER")
	}

	return json.Success(ctx, fiber.Map{"id": user.Id.String()})
}

func (s *UserService) Show(ctx *fiber.Ctx, id string) error {
	data, err := s.repository.FindById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return json.ErrorNotFound(ctx, err)
	}
	if err != nil {
		return json.Error(ctx, err, "ERR_SHOW_USER")
	}

	return json.Success(ctx, responses.UserDetailTransformer(data))
}

func (s *UserService) GetMe(ctx *fiber.Ctx) error {
	id := ctx.Locals("user_auth").(string)
	user, err := s.repository.FindById(id)
	if err != nil {
		return json.Error(ctx, err, "ERR_GET_ME_USER")
	}

	return json.Success(ctx, responses.UserDetailTransformer(user))
}

func (s *UserService) List(ctx *fiber.Ctx, paginate pagination.Pagination) error {
	paginationData, err := s.repository.GetAll(paginate)
	if err != nil {
		return json.Error(ctx, err, "ERR_LIST_USER")
	}

	return json.Pagination(ctx, paginationData)
}

func (s *UserService) Delete(ctx *fiber.Ctx, id string) error {
	err := s.repository.Delete(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return json.ErrorNotFound(ctx, err)
	}
	if err != nil {
		return json.Error(ctx, err, "ERR_DELETE_USER")
	}
	return json.Success(ctx, fiber.Map{"id": id})
}

func (s *UserService) generateEmailVerificationToken(email string) (string, error) {
	claims := middlewares.JwtCustomClaims{
		Issuer: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.GetJWTSecret())
}
