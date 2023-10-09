package controllers

import (
    "github.com/gofiber/fiber/v2"
)

type Endpoint struct {
    Name    string `json:"name"`
    URL     string `json:"url"`
    Example string `json:"example"`
}

type Response struct {
    Message   string     `json:"message"`
    Endpoints []Endpoint `json:"endpoints"`
    Debug     Debug      `json:"debug"`
}

type Debug struct {
    Error bool `json:"error"`
    Code  int  `json:"code"`
}

func NewIndexController(c *fiber.Ctx) error {
    endpoints := []Endpoint{
        {
            Name:    "Prefix Search",
            URL:     "/api/prefix?q={prefix}",
            Example: "/api/prefix?q=23.142.248.0/24",
        },
        {
            Name:    "ASN Search",
            URL:     "/api/asn?q={asn}",
            Example: "/api/asn?q=393577",
        },
    }

    response := Response{
        Message: "Hello! Welcome to rpki.online, a simple RPKI prefix validator API.",
        Endpoints: endpoints,
        Debug: Debug{
            Error: false,
            Code:  200,
        },
    }

    return c.Status(200).JSON(response)
}
