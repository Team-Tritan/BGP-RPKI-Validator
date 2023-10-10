package controllers

import (
    "encoding/json"
    "github.com/gofiber/fiber/v2"
    "fmt"
)

type APISearchResponse struct {
    ValidatedRoute ValidatedRouteData `json:"validated_route"`
}

type ValidatedRouteData struct {
    Route    RouteData    `json:"route"`
    Validity ValidityData `json:"validity"`
}

type RouteData struct {
    Prefix string `json:"prefix"`
    ASN    string `json:"origin_asn"`
}

type ValidityData struct {
    State       string `json:"state"`
    Description string `json:"description"`
}

func RPKISearchController(c *fiber.Ctx) error {
    prefix := c.Query("q")
    asn := c.Query("as")

    if prefix == "" || asn == "" {
        return c.Status(400).JSON(map[string]interface{}{
            "message": "Please provide a prefix and ASN.",
            "debug": map[string]interface{}{
                "error": true,
                "code": 400,
            },
        })
    }

    url := fmt.Sprintf("https://rpki-validator.ripe.net/api/v1/validity/%s/%s", asn, prefix)
    responseBody, statusCode, err := makeAPIRequest(url)

    if err != nil {
        return c.Status(500).JSON(map[string]interface{}{
            "message": fmt.Sprintf("Error making API request: %v", err),
            "debug": map[string]interface{}{
                "error": true,
                "code": 500,
            },
        })
    }

    var responseStruct APISearchResponse
    if err := json.Unmarshal(responseBody, &responseStruct); err != nil {
        return c.Status(500).JSON(map[string]interface{}{
            "message": fmt.Sprintf("Error unmarshaling API response: %v", err),
            "debug": map[string]interface{}{
                "error": true,
                "code": 500,
            },
        })
    }

    return c.Status(statusCode).JSON(responseStruct.ValidatedRoute)
}