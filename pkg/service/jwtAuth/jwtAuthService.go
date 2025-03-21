package jwtAuth

import (
	"doorProject/internal/domain/models"
	"doorProject/pkg/config"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthService struct {
	accessTokenCookieName      string
	refreshTokenCookieName     string
	jwtSecret                  string
	jwtRefreshSecret           string
	AccessTokenExpirationTime  int
	RefreshTokenExpirationTime int
}

func NewJWTAuthService(
	accessTokenCookieName string,
	refreshTokenCookieName string,
	jwtSecret string,
	jwtRefreshSecret string,
	tokenExpirationTime int,
) *AuthService {
	return &AuthService{
		accessTokenCookieName:     accessTokenCookieName,
		refreshTokenCookieName:    refreshTokenCookieName,
		jwtSecret:                 jwtSecret,
		jwtRefreshSecret:          jwtRefreshSecret,
		AccessTokenExpirationTime: tokenExpirationTime,
	}
}

func (a *AuthService) GenerateTokenAndSetCookie(worker *models.Worker, c echo.Context) error {
	accessToken, exp, err := a.generateAccessToken(worker)
	if err != nil {
		return err
	}

	a.setTokenCookie(accessToken, exp, c)

	refreshToken, exp, err := a.generateRefreshToken(worker)
	if err != nil {
		return err
	}
	a.setTokenCookie(refreshToken, exp, c)

	return nil
}

// Нужен для редиректа

func JWTErrorChecker(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("SignIn"))
}

func (a *AuthService) TokenRefresherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("user") == nil {
			return next(c)
		}

		u := c.Get("user").(*jwt.Token)
		claims := u.Claims.(*config.Claims)

		if time.Unix(claims.ExpiresAt.Unix(), 0).Sub(time.Now()) < 15*time.Minute {
			rc, err := c.Cookie(a.accessTokenCookieName)
			if err != nil {
				return jwt.ErrSignatureInvalid
			}

			if rc != nil && err == nil {
				tkn, err := jwt.ParseWithClaims(
					rc.Value, &config.Claims{},
					func(token *jwt.Token) (interface{}, error) {
						return []byte(a.jwtRefreshSecret), nil
					},
				)

				if err != nil {
					return jwt.ErrSignatureInvalid
				}

				if tkn != nil && tkn.Valid {
					_ = a.GenerateTokenAndSetCookie(
						&models.Worker{Name: claims.Name}, c,
					)
				}
			}
		}

		return next(c)
	}
}

func (a *AuthService) generateAccessToken(worker *models.Worker) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Duration(a.AccessTokenExpirationTime) * time.Minute)

	return a.generateToken(worker, expirationTime, []byte(a.jwtSecret))
}

func (a *AuthService) setTokenCookie(
	accessToken string,
	exp time.Time,
	c echo.Context,
) {
	cookie := new(http.Cookie)
	cookie.Name = a.accessTokenCookieName
	cookie.Value = accessToken
	cookie.Expires = exp
	cookie.Path = "/"
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}

func (a *AuthService) generateToken(
	worker *models.Worker,
	expirationTime time.Time,
	secret []byte,
) (string, time.Time, error) {
	claims := a.getClaims(worker, expirationTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

func (a *AuthService) getClaims(worker *models.Worker, expirationTime time.Time) config.Claims {
	claims := config.Claims{
		Name: worker.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: expirationTime,
			},
		},
	}
	return claims
}

func (a *AuthService) generateRefreshToken(worker *models.Worker) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Duration(a.RefreshTokenExpirationTime) * time.Hour)

	return a.generateToken(worker, expirationTime, []byte(a.jwtRefreshSecret))
}
