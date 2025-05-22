package jwt

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Documentar valor por defecto
const defaultExpirationTimeToken = 30

func GenerateToken(email string, logger logrus.FieldLogger) (string, string, error) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		logrus.Warn("Layer: jwt ", "Error al cargar el archivo .env:", err)
	}

	key := os.Getenv("SECRET_KEY")
	secretKey := []byte(key)

	expirationTimeStr := os.Getenv("TIME_TOKEN")
	expirationTimeDuration, err := strconv.Atoi(expirationTimeStr)

	if err != nil {
		logger.Errorln("Layer: Jwt", "Method: GenerateToken", "Error:", err)
		expirationTimeDuration = defaultExpirationTimeToken
	}

	refreshExpirationTimeDuration := 24 * time.Hour
	expirationTime := time.Now().Add(time.Duration(expirationTimeDuration) * time.Minute)
	refreshExpirationTime := time.Now().Add(refreshExpirationTimeDuration)

	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   email,
	}

	refreshClaims := &jwt.StandardClaims{
		ExpiresAt: refreshExpirationTime.Unix(),
		Subject:   email,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}
func ValidateToken(tokenStr string) (*jwt.StandardClaims, error) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		logrus.Warn("Layer: jwt ", "Error al cargar el archivo .env:", err)
	}
	key := os.Getenv("SECRET_KEY")
	secretKey := []byte(key)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, err
	}
	return claims, nil
}
