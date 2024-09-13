package helper

import (
	"errors"
	"fmt"
	"go-ecommerce-app/internal/domain"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)


type Auth struct {
	Secret string
}

func SetupAuth(s string) Auth {
	return Auth{
		Secret: s,
	}
}

func (a Auth) CreateHashedPassword(p string) (string, error) {
	
	if len(p) < 6 {
		return "", errors.New("password must be at least 6 characters")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	return string(hashedPassword), nil
}

func (a Auth) GenerateToken(id uint, email string, role string) (string, error) {

	if id == 0 || email == "" || role == "" {
		return "", errors.New("id, email, and role are required")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email": email,
		"role": role,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(a.Secret))
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return tokenString, nil
}

func (a Auth) VerifyPassword(plainPassword string, hashedPassword string) error {

	if len(plainPassword) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return errors.New("invalid password")
	}

	return nil
}

func (a Auth) VerifyToken(t string) (domain.User, error) {

	tokenArr := strings.Split(t, " ")
	if len(tokenArr) != 2 {
		return domain.User{}, errors.New("invalid token")
	}
	if tokenArr[0] != "Bearer" {
		return domain.User{}, errors.New("invalid token")
	}

	tokenStr := tokenArr[1]

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unknown signing method: %v", token.Header)
		}
		return []byte(a.Secret), nil
	})
	if err != nil {
		return domain.User{}, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return domain.User{}, errors.New("token expired")
		}

		user := domain.User{}
		user.ID = uint(claims["user_id"].(float64))
		user.Email = claims["email"].(string)
		user.UserType = claims["role"].(string)

		return user, nil
	}

	return domain.User{}, errors.New("token verification failed")
}

func (a Auth) Authorize(ctx *fiber.Ctx) error {

	authHeader := ctx.GetReqHeaders()["Authorization"]
	if len(authHeader) < 1 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
			"reason": "Authorization header is required",
		})
	}
	user, err := a.VerifyToken(authHeader[0])
	if err == nil && user.ID > 0 {
		ctx.Locals("user", user)
		return ctx.Next()
	} else {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
			"reason": err.Error(),
		})
	}

}

func (a Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {
	user := ctx.Locals("user")
	return user.(domain.User)
}

func (a Auth) GenerateCode() (int, error) {
	return RandomNumbers(6)
}

