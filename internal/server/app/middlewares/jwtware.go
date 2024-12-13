package middlewares

import (
	"github.com/Xurliman/auth-service/internal/constants"
	"github.com/Xurliman/auth-service/internal/server/app/repositories"
	"github.com/Xurliman/auth-service/pkg/json"
	"github.com/Xurliman/auth-service/pkg/log"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"os"
	"time"
)

type JwtCustomClaims struct {
	Issuer string `json:"issuer"`
	jwt.StandardClaims
}

type SkipperRoutesData struct {
	Method  string
	UrlPath string
}

func JwtMiddleware(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Set("X-XSS-Protection", "1; mode=block")
		ctx.Set("Strict-Transport-Security", "max-age=5184000")
		ctx.Set("X-DNS-Prefetch-Control", "off")

		for _, whiteList := range whiteListRoutes() {
			if ctx.Method() == whiteList.Method && ctx.Path() == whiteList.UrlPath {
				return ctx.Next()
			}
		}

		authorizationToken := ctx.Cookies(constants.SessionCookieName)
		if authorizationToken == "" {
			return json.ErrorUnauthorized(ctx, constants.ErrSessionNotFound)
		}

		jwtToken, err := jwt.ParseWithClaims(authorizationToken, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return json.ErrorUnauthorized(ctx, err)
		}
		claimsData := jwtToken.Claims.(*JwtCustomClaims)

		if !jwtToken.Valid {
			return json.ErrorUnauthorized(ctx, constants.ErrInvalidToken)
		}

		if claimsData.ExpiresAt < time.Now().Unix() {
			return json.ErrorUnauthorized(ctx, constants.ErrInvalidToken)
		}

		userRepo := repositories.NewAuthRepository(db)
		session, err := userRepo.FindSessionByToken(authorizationToken)
		if !session.IsActive {
			return json.ErrorUnauthorized(ctx, constants.ErrInvalidToken)
		}

		if err != nil {
			return json.ErrorUnauthorized(ctx, err)
		}
		log.Info("✅ SET USER AUTH")

		ctx.Locals("user_auth", claimsData.Issuer)
		ctx.Locals("user_session", session.Id.String())
		return ctx.Next()
	}
}

func whiteListRoutes() []SkipperRoutesData {
	return []SkipperRoutesData{
		{"POST", "/api/auth/login"},
		{"POST", "/api/auth/register"},
	}
}
