package main

import (
    "fmt"
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"

    "tritan.gg/rpki-validator/config"
    "tritan.gg/rpki-validator/middleware"
    "tritan.gg/rpki-validator/router"
)

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v", r)
        }
    }()

    app := buildApp()
    
    startServer(app)
}

func buildApp() *fiber.App {
    app := fiber.New()
    
    app.Use(logger.New())
    app.Use(cors.New())
    app.Use(middleware.CustomHeaders())

    return app
}

func startServer(app *fiber.App) {
    if err := router.NewRouter(app); err != nil {
        log.Fatalf("Error setting up routes: %v", err)
    }

    appConfig := config.AppConfigInstance
    if err := app.Listen(appConfig.HTTPPort); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}

