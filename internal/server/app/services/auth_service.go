package services

import (
	"github.com/Xurliman/auth-service/internal/constants"
	"github.com/Xurliman/auth-service/internal/server/app/interfaces"
	"github.com/Xurliman/auth-service/internal/server/app/middlewares"
	"github.com/Xurliman/auth-service/internal/server/app/models"
	"github.com/Xurliman/auth-service/internal/server/app/requests"
	"github.com/Xurliman/auth-service/internal/server/app/responses"
	"github.com/Xurliman/auth-service/pkg/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type AuthService struct {
	repository interfaces.IAuthRepository
}

func NewAuthService(repository interfaces.IAuthRepository) interfaces.IAuthService {
	return &AuthService{
		repository: repository,
	}
}

func (s *AuthService) Login(ctx *fiber.Ctx, request *requests.LoginRequest) error {
	if err := request.Validate(); err != nil {
		return json.ErrorValidation(ctx, err)
	}

	user, err := s.repository.FindByEmail(request.Email)
	if err != nil {
		return json.Error(ctx, err, "ERR_LOGIN")
	}

	if s.checkPasswordHash(request.Password, user.Password) {
		return json.ErrorUnauthorized(ctx, constants.ErrInvalidAuth)
	}

	expireHour, _ := time.ParseDuration(os.Getenv("JWT_EXPIRES") + "h")
	expiresAt := time.Now().Add(expireHour)
	token, err := s.generateToken(user.Id.String(), expiresAt.Unix())
	if err != nil {
		return json.ErrorUnauthorized(ctx, err)
	}

	_, err = s.repository.AddSession(models.UserSession{
		UserId:       user.Id,
		SessionToken: token,
		IsActive:     true,
		ExpiresAt:    expiresAt,
	})
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:        constants.SessionCookieName,
		Value:       token,
		Path:        "/",
		Expires:     expiresAt,
		Secure:      true,
		HTTPOnly:    true,
		SameSite:    "Strict",
		SessionOnly: false,
	})

	return json.Success(ctx, responses.AuthLoginTransformer(user, expiresAt.Unix()))
}

func (s *AuthService) Logout(ctx *fiber.Ctx) error {
	byToken, err := s.repository.FindSessionByToken(ctx.Cookies("user_session"))
	if err != nil {
		return json.Error(ctx, err, "E_LOGOUT_FAILED")
	}

	err = s.repository.MakeSessionInactive(byToken.Id.String())
	if err != nil {
		return json.Error(ctx, err, "E_LOGOUT_FAILED")
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     constants.SessionCookieName,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		Path:     "/",
		Secure:   true,
		HTTPOnly: true,
	})

	return json.Success(ctx, nil)
}

func (s *AuthService) generateToken(userGUID string, expiresAt int64) (string, error) {
	claims := middlewares.JwtCustomClaims{
		Issuer: userGUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func (s *AuthService) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
