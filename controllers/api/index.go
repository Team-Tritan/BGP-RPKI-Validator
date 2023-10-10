package controllers

import (
    "github.com/gofiber/fiber/v2"

    "io/ioutil"
	"net/http"
)

type Endpoint struct {
    Name        string `json:"name"`
    Description string `json:"description"`
    URL         string `json:"url"`
    Example     string `json:"example"`
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

func makeAPIRequest(url string) ([]byte, int, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, 500, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, 500, err
    }

    return body, resp.StatusCode, nil
}

func ApiIndexController(c *fiber.Ctx) error {
    endpoints := []Endpoint{
        {
            Name:    "RPKI Prefix Search",
            Description: "Searches for a prefix in the RPKI Validator.",
            URL:     "/api/rpki?q={prefix}&as={asn}",
            Example: "/api/rpki?q=23.142.248.0/24&as=393577",
        },
        {
            Name:    "ASN Prefix Search",
            Description: "Returns all prefixes for a given ASN.",
            URL:     "/api/prefixes?q={asn}",
            Example: "/api/prefixes?q=393577",
        },
    }

    response := Response{
        Message: "Hello! Welcome to our bgp toolkit API!",
        Endpoints: endpoints,
        Debug: Debug{
            Error: false,
            Code:  200,
        },
    }

    return c.Status(200).JSON(response)
}
