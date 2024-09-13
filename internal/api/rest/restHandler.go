package rest


import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"go-ecommerce-app/internal/helper"
)

type RestHandler struct {
	App		*fiber.App
	DB		*gorm.DB
	Auth	helper.Auth
}