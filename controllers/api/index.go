package controllers

import (
    "github.com/gofiber/fiber/v2"

    "io/ioutil"
	"net/http"
)

type Endpoint struct {
    Name            string `json:"name"`
    Description     string `json:"description"`
    URLQueries string `json:"url_queries"`
    URL             string `json:"url"`
    Example         string `json:"example"`
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
            Description: "Searches for a specific prefix or all ASN prefixes in the RPKI Validator.",
            URLQueries: "as: AS Number",
            URL:     "/api/rpki?q={prefix}&as={asn}",
            Example: "/api/rpki?as=393577",
        },
        {
            Name:    "ASN Prefix Search",
            Description: "Returns all prefixes for a given ASN.",
            URLQueries: "q: AS Number",
            URL:     "/api/prefixes?q={asn}",
            Example: "/api/prefixes?q=393577",

        },
    }

    response := Response{
        Message: "Welcome to our rpki.online, a bgp toolkit API!",
        Endpoints: endpoints,
        Debug: Debug{
            Error: false,
            Code:  200,
        },
    }

    return c.Status(200).JSON(response)
}
