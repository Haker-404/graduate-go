package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SecretKey = "JingwenLi"

type jwtCustomClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	Seq      string
	Username string
	Admin    bool
}

func CreateToken(seq string, username string, isAdmin bool) (tokenString string, err error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 1).Unix()),
		},
		seq,
		username,
		isAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(SecretKey))
	return
}
func ParseToken(tokenSrt string, SecretKey []byte) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claims = token.Claims
	return
}
func AuthCheck(tokenString string) bool {
	//[7:]去掉Bearer前缀
	token, err := jwt.ParseWithClaims(tokenString[7:], &jwtCustomClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err = token.Claims.Valid(); err != nil {
		fmt.Print(err)
		return false
	} else {
		claims := token.Claims.(*jwtCustomClaims)
		fmt.Println(claims.Username)
		return true
	}
}
