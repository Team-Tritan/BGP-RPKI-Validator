package router

import (
    "github.com/gofiber/fiber/v2"

    "tritan.gg/rpki-validator/controllers/api"
)

func NewRouter(app *fiber.App) error {   
    // UI \\ 
    // TODO: Add UI Templating
    ui := app.Group("/")
    ui.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    // API \\
    // TODO: Add API versioning, asn and prefix fetch
    apiGroup := app.Group("/api")
    apiGroup.Get("/", controllers.ApiIndexController)

    return nil
}
