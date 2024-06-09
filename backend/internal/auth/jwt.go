package auth

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/quangd42/meal-planner/backend/internal/database"
)

const (
	JWTIssuer                 = "meal_planner"
	DefaultExpirationDuration = time.Hour * 24
)

var (
	ErrTokenNotFound    = errors.New("token not found")
	ErrClaimTypeInvalid = errors.New("claim type cannot be verified")
)

var jwtSecret string

func init() {
	godotenv.Load()
	jwtSecret = os.Getenv("JWT_SECRET")
}

type UserClaims struct {
	UserID uuid.UUID `json:"userID"`
	jwt.RegisteredClaims
}

func (uc *UserClaims) GetUserID() uuid.UUID {
	return uc.UserID
}

func CreateJWT(user database.User, d time.Duration) (string, error) {
	now := time.Now().UTC()
	claims := UserClaims{
		user.ID,
		jwt.RegisteredClaims{
			Issuer:    JWTIssuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(d)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func VerifyJWT(tokenString string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return uuid.UUID{}, ErrClaimTypeInvalid
	}

	return claims.UserID, nil
}

func GetHeaderToken(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if len(header) < 7 && strings.ToLower(header[0:6]) != "bearer" {
		return "", ErrTokenNotFound
	}
	return header[7:], nil
}
