package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	ID    string `json:"userId"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

func main() {

	customClaims := make(jwt.MapClaims)
	customClaims["name"] = "CrazyWolf"
	customClaims["id"] = 1
	t := time.Now().UnixNano() / 1e6
	customClaims["time"] = fmt.Sprintf("%v", t)

	key := "woshijiamimiyao1"

	token, err := getToken(key, customClaims)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)

	c, b := parseToken(token, key)
	fmt.Println(b)
	fmt.Println(c)
}

func parseToken(tokenString string, key string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(jwttoken *jwt.Token) (interface{}, error) {
		if _, ok := jwttoken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", jwttoken.Header["alg"])
		}
		return []byte(key), nil
	})

	fmt.Println(token.Claims)
	fmt.Println(token.Valid)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return "", false
	}

}

func getToken(key string, customClaims jwt.Claims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = customClaims

	res, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return res, nil
}
