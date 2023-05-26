package token

import (
	"fmt"
	"project_blog_gin/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var API_SECRET = utils.GetEnv("API_SECRET", "rahasia")

func GenerateToken(user_id uint) (string, error){
	token_hour, err := strconv.Atoi(utils.GetEnv("TOKEN_HOUR", "1")) // ubah str ke integer
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_hour)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(API_SECRET))
}

func ExtractToken(c *gin.Context) string{
	token := c.Query("token")

	if token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("authorization")
	if len(strings.Split(bearerToken, "")) == 2 {
		return strings.Split(bearerToken, "")[1]
	}
	return ""
}

func TokenValid(c *gin.Context) error{
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC);
		!ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}		
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractTokenID(c *gin.Context) (uint, error){
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC);
		!ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}		
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32) //size integer 10, bit 32
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}

