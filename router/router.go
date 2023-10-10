package router

import (
    "github.com/gofiber/fiber/v2"

    "tritan.gg/rpki-validator/controllers/api"
)

func NewRouter(app *fiber.App) error {   
    // TODO: Add UI Templating
    ui := app.Group("/")
    ui.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    apiGroup := app.Group("/api")
    apiGroup.Get("/", controllers.ApiIndexController)
    apiGroup.Get("/rpki", controllers.RPKISearchController)
    apiGroup.Get("/prefixes", controllers.PrefixSearchController)

    return nil
}
