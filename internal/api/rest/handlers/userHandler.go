package handlers

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"

	"github.com/gofiber/fiber/v2"
)


type UserHandler struct {
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := service.UserService{
		Repo: repository.NewUserRepository(rh.DB),
		Auth: rh.Auth,
		Config: rh.Config,
	}
	handler := UserHandler{
		svc: svc,
	}

	// public routes
	publicRoutes := app.Group("/users")

	publicRoutes.Post("/register", handler.Register)
	publicRoutes.Post("/login", handler.Login)

	// private routes
	privateRoutes := publicRoutes.Group("/", rh.Auth.Authorize)

	privateRoutes.Get("/profile", handler.GetProfile)
	privateRoutes.Get("/verify", handler.GetVerificationCode)
	privateRoutes.Post("/verify", handler.Verify)
	privateRoutes.Post("/profile", handler.CreateProfile)

	privateRoutes.Post("/cart", handler.AddToCart)
	privateRoutes.Get("/cart", handler.GetCart)
	privateRoutes.Get("/order", handler.GetOrders)
	privateRoutes.Get("/order/:id", handler.GetOrder)

	privateRoutes.Get("/become-seller", handler.BecomeSeller)
}

func (uh *UserHandler) Register(ctx *fiber.Ctx) error {
	user :=  dto.UserRegisterDto{}
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid input",
		})
	}

	token, err := uh.svc.Register(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Registered",
		"token": token,
	})
}

func (uh *UserHandler) Login(ctx *fiber.Ctx) error {

	loginInput :=  dto.UserLoginDto{}
	err := ctx.BodyParser(&loginInput)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid input",
		})
	}

	token, err := uh.svc.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"message": "Invalid credientials",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Login",
		"token": token,
	})
}

func (uh *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	user := uh.svc.Auth.GetCurrentUser(ctx)
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Get Profile",
		"user": user,
	})
}

func (uh *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {

	user := uh.svc.Auth.GetCurrentUser(ctx)

	err := uh.svc.GetVerificationCode(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "get verification code",
	})
}

func (uh *UserHandler) Verify(ctx *fiber.Ctx) error {

	user := uh.svc.Auth.GetCurrentUser(ctx)

	var req dto.VerificationCodeDto
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid input",
		})
	}

	err := uh.svc.VerifyCode(user.ID, req.Code)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user verified successfully",
	})
}
func (uh *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login",
	})
}

func (uh *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login",
	})
}

func (uh *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login",
	})
}

func (uh *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login",
	})
}

func (uh *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login",
	})
}

func (uh *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login",
	})
}