package helper

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rahmaninsani/dipay-backend-technical-test/02-restful-api/employee-service/model/domain"
	"time"
)

type Claims struct {
	TokenID  uuid.UUID `json:"token_id"`
	Username string    `json:"username"`
	*jwt.RegisteredClaims
}

func GenerateToken(admin *domain.Admin, ttl time.Duration, secretKey string) (string, error) {
	claims := &Claims{
		TokenID:  uuid.New(),
		Username: admin.Username,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(ttl)),
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	
	return signedToken, nil
}

func ValidateToken(encodedToken string, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (any, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("signing method invalid")
		}
		
		return []byte(secretKey), nil
	})
	
	if err != nil {
		return nil, err
	}
	
	tokenClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	
	return tokenClaims, nil
}
