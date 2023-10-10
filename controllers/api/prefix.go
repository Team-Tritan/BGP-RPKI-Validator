package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"

)

type ASNData struct {
	Type  string   `json:"type"`
	ASNs  []string `json:"asns"`
	Meta  interface{} `json:"meta"`
	Result struct {
		Relations []struct {
			Type    string `json:"type"`
			Members []struct {
				Prefix string `json:"prefix"`
				Meta   []struct {
					SourceType string   `json:"sourceType"`
					SourceID   string   `json:"sourceID"`
					OriginASNs []string `json:"originASNs"`
				} `json:"meta"`
			} `json:"members"`
		} `json:"relations"`
	} `json:"result"`
}

type ASNSearchResponse struct {
	Prefix  string   `json:"prefix"`
	Meta    []struct {
		SourceType string   `json:"sourceType"`
		SourceID   string   `json:"sourceID"`
		OriginASNs []string `json:"originASNs"`
	} `json:"meta"`
}

func PrefixSearchController(c *fiber.Ctx) error {
	asn := c.Query("q")

	if asn == "" {
		return c.Status(400).JSON(map[string]interface{}{
			"message": "Please provide an ASN.",
			"debug": map[string]interface{}{
				"error": true,
				"code": 400,
			},
		})
	}

	url := fmt.Sprintf("https://rest.bgp-api.net/api/v1/asn/%s/search", asn)
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

	var asnData ASNData
	if err := json.Unmarshal(responseBody, &asnData); err != nil {
		return c.Status(500).JSON(map[string]interface{}{
			"message": fmt.Sprintf("Error unmarshaling API response: %v", err),
			"debug": map[string]interface{}{
				"error": true,
				"code": 500,
			},
		})
	}

	var response []ASNSearchResponse
	for _, relation := range asnData.Result.Relations {
		for _, member := range relation.Members {
			response = append(response, ASNSearchResponse{
				Prefix: member.Prefix,
				Meta:   member.Meta,
			})
		}
	}

	return c.Status(statusCode).JSON(response)
}
