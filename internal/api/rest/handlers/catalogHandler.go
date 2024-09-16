package handlers

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
)


type CatalogHandler struct {
	svc service.CatalogService
}

func SetupCatalogRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc:= service.CatalogService{
		Repo: repository.NewCatalogRepository(rh.DB),
		Auth: rh.Auth,
		Config: rh.Config,
	}

	handler := CatalogHandler{
		svc: svc,
	}

	// public routes for listing products and categories
	app.Get("/products")
	app.Get("/products/:id")
	app.Get("/categories")
	app.Get("/categories/:id")

	// private routes for creating, updating, and deleting products and categories
	sellerRoutes := app.Group("/seller", rh.Auth.AuthorizeSeller)

	// categories
	sellerRoutes.Post("/categories", handler.CreateCategory)
	sellerRoutes.Patch("/categories/:id", handler.UpdateCategory)
	sellerRoutes.Delete("/categories/:id", handler.DeleteCategory)

	// products
	sellerRoutes.Post("/products", handler.CreateProduct)
	sellerRoutes.Get("/products", handler.GetProducts)
	sellerRoutes.Get("/products/:id", handler.GetProduct)
	sellerRoutes.Put("/products/:id", handler.UpdateProduct)
	sellerRoutes.Patch("/products/:id", handler.UpdateStock)
	sellerRoutes.Delete("/products/:id", handler.DeleteProduct)

}

func (h CatalogHandler) CreateCategory(ctx *fiber.Ctx) error {

	user := h.svc.Auth.GetCurrentUser(ctx)
	log.Printf("Current user: %v", user.ID)

	return rest.SuccessResponse(ctx, "Category created successfully", nil)
}

func (h CatalogHandler) UpdateCategory(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "Category updated successfully", nil)
}

func (h CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "Category deleted successfully", nil)
}


func (h CatalogHandler) CreateProduct(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "Product created successfully", nil)
}

func (h CatalogHandler) GetProducts(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "Products retrieved successfully", nil)
}

func (h CatalogHandler) GetProduct(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "Product retrieved successfully", nil)
}

func (h CatalogHandler) UpdateProduct(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "Product updated successfully", nil)
}

func (h CatalogHandler) UpdateStock(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "Stock updated successfully", nil)
}

func (h CatalogHandler) DeleteProduct(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "Product deleted successfully", nil)
}
