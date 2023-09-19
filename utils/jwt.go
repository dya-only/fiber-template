package utils

//
//import (
//	"github.com/golang-jwt/jwt/v4"
//	"os"
//	"time"
//)
//
//func NewToken(id int64) string {
//	claims := jwt.MapClaims{
//		"id":  id,
//		"exp": time.Now().Add(time.Hour * 72).Unix(),
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
//	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
//}
