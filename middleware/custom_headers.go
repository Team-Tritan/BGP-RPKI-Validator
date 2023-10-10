package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CustomHeaders() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		elapsedMilliseconds := time.Since(start).Milliseconds()
		responseTime := fmt.Sprintf("%.1fms", float64(elapsedMilliseconds))

		c.Set("X-Response-Time", responseTime)
		c.Set("X-Developers", "Tritan Devs (AS393577)")
		c.Set("X-Hello", "Why are you looking at the headers? Wanna join a team of devs and make cool stuff? https://discord.gg/http")

		if err != nil {
			log.Printf("Middleware error: %v", err)
			return err
		}

		return nil
	}
}
