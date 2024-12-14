package services

import (
	"github.com/Xurliman/auth-service/internal/config/config"
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
	"strconv"
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

	if !user.IsEmailVerified {
		return json.Error(ctx, constants.ErrEmailNotVerified, "ERR_EMAIL_NOT_VERIFIED")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return json.ErrorUnauthorized(ctx, constants.ErrInvalidAuth)
	}

	cfg := config.Setup()
	expireHour, err := time.ParseDuration(strconv.Itoa(cfg.JWT.Expires) + "h")
	if err != nil {
		return json.Error(ctx, err, "ERR_LOGIN")
	}

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

	isSecure := cfg.App.Env == "production"
	ctx.Cookie(&fiber.Cookie{
		Name:        constants.SessionCookieName,
		Value:       token,
		Path:        "/",
		Expires:     expiresAt,
		Secure:      isSecure,
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.GetJWTSecret())
}

func (s *AuthService) VerifyEmail(ctx *fiber.Ctx, tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &middlewares.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.GetJWTSecret(), nil
	})
	if err != nil {
		return json.Error(ctx, constants.ErrInvalidToken, "ERR_INVALID_TOKEN")
	}

	claims, ok := token.Claims.(*middlewares.JwtCustomClaims)
	if !ok || !token.Valid || claims.ExpiresAt < time.Now().Unix() {
		return json.ErrorUnauthorized(ctx, constants.ErrInvalidToken)
	}

	err = s.repository.MakeEmailVerified(claims.Issuer)
	if err != nil {
		return json.Error(ctx, err, "ERR_VERIFY_EMAIL")
	}

	return json.Success(ctx, "Now you can login")
}
