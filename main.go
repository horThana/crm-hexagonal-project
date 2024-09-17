package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/horThana/Backend/adapters/http"
	"github.com/horThana/Backend/adapters/repository"
	"github.com/horThana/Backend/core/domain"
	"github.com/horThana/Backend/core/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
    app := fiber.New()

    // Connect to product database
    productDB, err := gorm.Open(sqlite.Open("product.db"), &gorm.Config{})
    if err != nil {
        panic("ไม่สามารถเชื่อมต่อฐานข้อมูลผลิตภัณฑ์ได้")
    }

    // Connect to user database
    userDB, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
    if err != nil {
        panic("ไม่สามารถเชื่อมต่อฐานข้อมูลผู้ใช้ได้")
    }

    // Migrate the schema for both databases
    productDB.AutoMigrate(&domain.Product{})
    userDB.AutoMigrate(&domain.User{})

    // Set up repositories, services, and handlers for products
    productRepo := repository.NewGormProductRepository(productDB)
    productService := services.NewProductService(productRepo)
    productHandler := http.NewHttpProductAdapter(productService)

    // Set up repositories, services, and handlers for users
    userRepo := repository.NewGormUserRepository(userDB)
    userService := services.NewUserService(userRepo)
    userHandler := http.NewHttpUserAdapter(userService)

    // Define routes for products
    app.Post("/product", productHandler.CreateProduct)
    app.Get("/product/:id", productHandler.FindProductByID)
    app.Get("/product", productHandler.FindAllProducts)
    app.Delete("/product/:id", productHandler.DeleteProduct)

    // Define routes for users
    app.Post("/user", userHandler.CreateUser)
    app.Get("/user/:id", userHandler.FindUserByID)
    app.Get("/user", userHandler.FindAllUsers)
    app.Delete("/user/:id", userHandler.DeleteUser)

    // Start the server
    app.Listen(":8000")
}