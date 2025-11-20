package utils

import (
"os"
"time"

"github.com/golang-jwt/jwt/v5"


)

var jwtKey []byte

func init() {
key := os.Getenv("JWT_SECRET")
if key == "" {
// fallback agar tidak crash
key = "DEFAULT_SECRET"
}
jwtKey = []byte(key)
}

func GenerateToken(userID int) (string, error) {
claims := jwt.MapClaims{
"user_id": userID,
"exp": time.Now().Add(24 * time.Hour).Unix(),
}

token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
return token.SignedString(jwtKey)


}

func ValidateToken(tokenStr string) (*jwt.Token, error) {
return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
return jwtKey, nil
})
}