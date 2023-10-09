package router

import (
    "github.com/gofiber/fiber/v2"
    "tritan.gg/rpki-validator/controllers"
)

func NewRouter(app *fiber.App) error {    
	app.Get("/", controllers.NewIndexController)

    return nil
}
