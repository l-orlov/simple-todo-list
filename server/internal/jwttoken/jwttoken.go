package jwttoken

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// tokenLifeTime - срок действия токена
const tokenLifeTime = time.Hour

// Секретный ключ для подписи токена (храните его в безопасности)
var secretKey = []byte("yourSecretKey")

// GenerateToken генерирует Bearer Token
func GenerateToken(userID string) (string, error) {
	now := time.Now()

	// Параметры токена (время жизни, алгоритм подписи и т.д.)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,                        // Сохраняем user id
		"exp": now.Add(tokenLifeTime).Unix(), // Токен будет действителен в течение указанного lifetime
		"iat": now,                           // Токен подписан в этом время
		"nbf": now.Unix(),                    // Токен может быть применен не раньше этого времени
	})

	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken проверяет токен
func ValidateToken(tokenStr string) (userID string, err error) {
	// Парсим токен
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Проверка используемого алгоритма
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	// Проверка ошибок при парсинге токена
	if err != nil {
		return "", fmt.Errorf("jwt.Parse: %w", err)
	}

	// Проверка валидности токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("token is invalid")
	}

	userID = claims["sub"].(string)

	return userID, nil
}
