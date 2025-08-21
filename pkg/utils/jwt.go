package utils

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	key []byte = []byte(os.Getenv("JWT_SECRET"))
	t   *jwt.Token
)

func SignJWTToken(userId int64, email string) (string, error) {
	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":   "workout-tracker",
			"sub":   userId,
			"email": email,
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		})
	s, err := t.SignedString(key)

	return s, err
}

func verifyJWTToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorExpired)
	}
	
	// Check if the token has expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && claims["exp"] != nil {
		if exp, ok := claims["exp"].(float64); ok && time.Unix(int64(exp), 0).Before(time.Now()) {
			return nil, jwt.NewValidationError("token has expired", jwt.ValidationErrorExpired)
		}
	}
	return token, nil
}

func getJWTTokenFromHeader(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return "", jwt.NewValidationError("authorization header is missing", jwt.ValidationErrorMalformed)
	}
	if len(token) < 7 || token[:7] != "Bearer " {
		return "", jwt.NewValidationError("invalid authorization header format", jwt.ValidationErrorMalformed)
	}
	token = token[7:]
	if token == "" {
		return "", jwt.NewValidationError("token is empty", jwt.ValidationErrorMalformed)
	}
	return token, nil
}

func ExtractUserIdFromJWTToken(r *http.Request) (int64, error) {
	tokenString, err := getJWTTokenFromHeader(r)
	if err != nil {
		return 0, err
	}

	token, err := verifyJWTToken(tokenString)
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && claims["sub"] != nil {
		if userId, ok := claims["sub"].(float64); ok {
			return int64(userId), nil
		}
	}
	return 0, jwt.NewValidationError("user ID not found in token", jwt.ValidationErrorClaimsInvalid)
}
