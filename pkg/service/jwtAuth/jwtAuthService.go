package jwtAuth

import (
	"doorProject/internal/domain/models"
	"doorProject/internal/repository"
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
	accessTokenExpirationTime  int
	refreshTokenExpirationTime int
	tokenRepo                  repository.RefreshTokenRepository
	workerRepo                 repository.WorkerRepositoryInterface
}

func NewJWTAuthService(
	accessTokenCookieName string,
	refreshTokenCookieName string,
	jwtSecret string,
	jwtRefreshSecret string,
	tokenExpirationTime int,
	refreshTokenExpirationTime int,
	r repository.RefreshTokenRepository,
	w repository.WorkerRepositoryInterface,
) *AuthService {
	return &AuthService{
		accessTokenCookieName:      accessTokenCookieName,
		refreshTokenCookieName:     refreshTokenCookieName,
		jwtSecret:                  jwtSecret,
		jwtRefreshSecret:           jwtRefreshSecret,
		accessTokenExpirationTime:  tokenExpirationTime,
		refreshTokenExpirationTime: refreshTokenExpirationTime,
		tokenRepo:                  r,
		workerRepo:                 w,
	}
}

func (a *AuthService) GenerateTokenAndSetCookie(worker *models.Worker, c echo.Context) (string, error) {
	accessToken, exp, err := a.generateAccessToken(worker)
	if err != nil {
		return "", err
	}

	a.setTokenCookie(a.accessTokenCookieName, accessToken, exp, c)

	refreshToken, exp, err := a.generateRefreshToken(worker)

	if err != nil {
		return "", err
	}

	refreshTokenModel := &models.UserToken{
		ExpiredAt: exp,
		WorkerID:  worker.ID,
		TokenHash: refreshToken,
		IsValid:   true,
	}

	err = a.tokenRepo.CreateRefreshToken(refreshTokenModel)
	if err != nil {
		return "", err
	}

	a.setTokenCookie(a.refreshTokenCookieName, refreshToken, exp, c)

	return accessToken, nil
}

func (a *AuthService) RefreshAccessToken(c echo.Context) (string, error) {
	rc, err := c.Cookie(a.refreshTokenCookieName)
	if err != nil {
		return "", err
	}

	token, err := a.tokenRepo.FindRefreshTokenByToken(rc.Value)
	if err != nil {
		return "", err
	}

	w, err := a.workerRepo.FindUserById(token.WorkerID)
	if err != nil {
		return "", err
	}

	err = a.tokenRepo.InvalidateRefreshToken(token)
	if err != nil {
		return "", err
	}

	accessToken, err := a.GenerateTokenAndSetCookie(w, c)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (a *AuthService) JWTErrorChecker(c echo.Context, err error) error {
	return c.JSON(
		http.StatusUnauthorized, map[string]string{
			"message": "Требуется авторизация",
		},
	)
}

func (a *AuthService) generateAccessToken(worker *models.Worker) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Duration(a.accessTokenExpirationTime) * time.Minute)

	return a.generateToken(worker, expirationTime, []byte(a.jwtSecret))
}

func (a *AuthService) setTokenCookie(
	accessTokenCookieName string,
	accessToken string,
	exp time.Time,
	c echo.Context,
) {
	cookie := new(http.Cookie)
	cookie.Name = accessTokenCookieName
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
	expirationTime := time.Now().Add(time.Duration(a.refreshTokenExpirationTime) * time.Hour)

	return a.generateToken(worker, expirationTime, []byte(a.jwtRefreshSecret))
}
