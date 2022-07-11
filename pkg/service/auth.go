package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	todo "github.com/Maksat-luci/REST-API-TODO-service"
	"github.com/Maksat-luci/REST-API-TODO-service/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "amaymon20021108"
	signingKey = "zhiznetostanovlenieluchoiversieysebya"
	tokenTTL   = 12 * time.Hour
)

// создаём расширенную версию структуры jwt.StandardClaims
// наследуемся от него же и дополняем его филдом UserID
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	// в бизнес логике пока что просто хешируем пароль пользователя
	// далее мы используем слой базы данных для сохранения данных в базе
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	// получаем юзера с базы данных
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	// генерируем новый токен и указываем когда он истечёт и когда был создан
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})
	// возвращаем подписанный токен
	return token.SignedString([]byte(signingKey))
}

//функция для хеширования пароля так как в базе данных, пароли не должны хранится в открытом в виде
func generatePasswordHash(password string) string {
	hash := sha1.New()
	// создаём обьект sha1 для того чтобы использовать методы для хеширования
	hash.Write([]byte(password))
	// используем фичу криптографии называемый солью, для лучшего хеширования пароля
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}


func (s *AuthService) ParseToken(accesToken string) (int, error) {
	//тут мы проверяем на метод подписи токена если она корректна то возвращаем подпись
	TokenUser, err := jwt.ParseWithClaims(accesToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, nil
	}
	// проверяем является ли нвш токенклеймс токен клеймсом
	claims, ok := TokenUser.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	// если всё хорошо то возвращаем user.id
	return claims.UserId, nil

}
