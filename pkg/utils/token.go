package utils

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func GenerateCodeAndTime() (string, int64) {
	jwt_time := os.Getenv("JWT_EXPIRATION_TIME")
	jwtDuration, _ := strconv.Atoi(jwt_time)
	code := make([]byte, 6)
	rand.Read(code)
	result := fmt.Sprintf("%x", code)
	exp := time.Now().Add(time.Hour * time.Duration(jwtDuration)).Unix()
	return result, exp
}
