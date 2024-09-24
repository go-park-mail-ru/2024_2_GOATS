package auth

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/model"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
	"unicode/utf8"
)

type ICached interface {
	GetToken(ctx context.Context, token string) (string, error)
	SetToken(ctx context.Context, token string, data string, expiration time.Duration) error
	Close() error
}

type IRepository interface {
	SaveUser(ctx context.Context, email string, passwordHash string, nickname string, sex string, birthdate time.Time) error
	GetUserByEmail(ctx context.Context, email string) (string, error) // you can add a full user model here
}

// AuthService implements the business logic for user authentication
type AuthService struct {
	repo  IRepository
	cache ICached
}

func NewAuthService(
	repo IRepository,
	cache ICached,
) *AuthService {
	return &AuthService{
		repo:  repo,
		cache: cache,
	}
}

func (s *AuthService) Register(ctx context.Context, email string, password string, passwordConfirm string, nickname string, sex string, birthdate *timestamppb.Timestamp) (bool, string, []*model.ErrorMessage) {
	var errors []*model.ErrorMessage

	if password != passwordConfirm {
		errors = append(errors, &model.ErrorMessage{Error: "Passwords do not match"})
		return false, "", errors
	}

	if utf8.RuneCountInString(password) < 6 {
		errors = append(errors, &model.ErrorMessage{Error: "Password too short"})
		return false, "", errors
	}

	birthTime := birthdate.AsTime()

	// Сохраняем юзера
	err := s.repo.SaveUser(ctx, email, hashPassword(password), nickname, sex, birthTime)
	if err != nil {
		errors = append(errors, &model.ErrorMessage{Error: "Failed to register user"})
		return false, "", errors
	}

	// Генерируем токен
	token, err := NewToken(email, time.Second*100) // Укажите продолжительность токена
	if err != nil {
		errors = append(errors, &model.ErrorMessage{Error: "Failed to create token"})
		return false, "", errors
	}

	// Сохраняем токен в кэше (если необходимо)
	err = s.cache.SetToken(ctx, token, email, time.Second*100)
	if err != nil {
		errors = append(errors, &model.ErrorMessage{Error: "Failed to save token"})
		return false, "", errors
	}

	return true, token, nil
}

func (s *AuthService) Login(ctx context.Context, email string, password string) (bool, string, []*model.ErrorMessage) {
	var errors []*model.ErrorMessage

	storedPasswordHash, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		println("Failed to get user by email")
		errors = append(errors, &model.ErrorMessage{Error: "User not found"})
		return false, "", errors
	}

	if !checkPasswordHash(password, storedPasswordHash) {
		errors = append(errors, &model.ErrorMessage{Error: "Invalid password"})
		return false, "", errors
	}

	// Генерим токен и сохраняем в кэш
	//token := uuid.New().String()
	token, err := NewToken(email, time.Second*100)
	if err != nil {
		errors = append(errors, &model.ErrorMessage{Error: "Failed to create token"})
	}
	//err = s.cache.SetToken(ctx, token, email, time.Hour*24)
	err = s.cache.SetToken(ctx, token, email, time.Second*100)
	if err != nil {
		fmt.Println(err)
		errors = append(errors, &model.ErrorMessage{Error: "Failed to create session"})
		return false, "", errors
	}

	//return true, nil
	return true, token, nil

}

func NewToken(email string, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	//claims := token.Claims.(jwt.MapClaims)
	return token.SignedString([]byte(email))
}

func (s *AuthService) Logout(ctx context.Context, email string) (bool, []*model.ErrorMessage) {
	var errors []*model.ErrorMessage
	fmt.Println("logout", email)
	// возможно сделать недействительным токен в кеше
	data, err := s.cache.GetToken(ctx, email)
	fmt.Println("data", data)
	err = s.cache.SetToken(ctx, data, "", time.Second*1)
	if err != nil {
		errors = append(errors, &model.ErrorMessage{Error: "Failed to logout"})
		return false, errors
	}

	err = s.cache.SetToken(ctx, email, "", time.Second*1)
	if err != nil {
		errors = append(errors, &model.ErrorMessage{Error: "Failed to logout"})
		return false, errors
	}

	return true, nil
}

func (s *AuthService) Close() error {
	return s.cache.Close()
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}
	return string(hashedPassword)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
