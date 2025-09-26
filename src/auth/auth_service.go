package auth

import (
	"cashier-be/pkg/models"
	"cashier-be/src/user"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(request *LoginRequest) (JwtToken, error)
	Profile(userId uint) (user.UserProfileData, error)
	Register(request *RegisterRequest) (JwtToken, error)
	GenerateToken(userId string) (string, error)
}

func HashPassword(password string) (string, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), nil
}

func checkPassword(hashedPassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}

type AuthService struct {
	UserRepository user.IUserRepository
}

func (service *AuthService) GenerateToken(userId string) (string, error) {
	secretKey := viper.Get("APP_SECRET_KEY").(string)
	secretKeyByte := []byte(secretKey)
	iss := fmt.Sprintf("%v:%v", "jwtservice", "3321")
	claims := jwt.MapClaims{
		"sub":   userId,
		"exp":   time.Now().Add(time.Hour * 3).Unix(),
		"iss":   iss,
		"roles": []string{"admin"},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["kid"] = "sim2"
	signedToken, _ := token.SignedString(secretKeyByte)

	return signedToken, nil

}

func (service *AuthService) Register(request *RegisterRequest) (JwtToken, error) {

	password, _ := HashPassword(request.Password)

	resUser, err := service.UserRepository.Create(user.UserData{
		Name:     request.Name,
		Username: request.Username,
		Role:     models.Role(request.Role),
		Password: password,
	})

	if err != nil {
		return JwtToken{}, err
	}

	resToken, _ := service.GenerateToken(fmt.Sprintf("%d", resUser.ID))

	return JwtToken{
		User: user.UserProfileData{
			Id:       resUser.ID,
			Role:     resUser.Role,
			Name:     resUser.Name,
			Username: resUser.Username,
		},
		Token: resToken,
	}, nil
}

func (service *AuthService) Login(request *LoginRequest) (JwtToken, error) {

	resUser, err := service.UserRepository.FindByUsername(request.Username)
	if err != nil {
		return JwtToken{}, errors.New("credential wrong")
	}

	if !checkPassword(resUser.Password, request.Password) {
		return JwtToken{}, errors.New("credential wrong")
	}

	resToken, _ := service.GenerateToken(fmt.Sprintf("%d", resUser.ID))
	return JwtToken{
		User: user.UserProfileData{
			Id:       resUser.ID,
			Name:     resUser.Name,
			Username: resUser.Username,
		},
		Token: resToken,
	}, nil
}

func (service *AuthService) Profile(userId uint) (user.UserProfileData, error) {

	resUser, err := service.UserRepository.FindById(userId)
	if err != nil {
		return user.UserProfileData{}, err
	}

	return user.UserProfileData{
		Id:       resUser.ID,
		Name:     resUser.Name,
		Username: resUser.Username,
	}, nil
}
