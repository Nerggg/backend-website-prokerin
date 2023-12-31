package utilities

import (
	"backend-prokerin/config"
	"backend-prokerin/models"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

var privateKey = []byte(config.Config.PrivateKey)
var tokenTTL = config.Config.TokenTTL

func GenerateJWT(user models.UserAccount) (string, error) {
	tokenTTL, _ := strconv.Atoi(tokenTTL)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
		"exp": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})

	return token.SignedString(privateKey)
}
func ValidateJWT(context *gin.Context) error {
	token, err := getToken(context)

	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return nil
	}
	log.Info().Msgf("Topic: %v", token)
	return errors.New("invalid token provided")
}

func getToken(context *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	log.Info().Msgf("Topic: %v", bearerToken)
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}

func ValidatePassword(password string, password_user string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(password_user))
	return err == nil
}
func CurrentUser(context *gin.Context) (string, error) {
	err := ValidateJWT(context)
	if err != nil {
		return "", err
	}
	token, _ := getToken(context)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := (claims["id"].(string))
	return userId, nil
}
