package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"tritan.gg/rpki-validator/config"
)

type RouteValidity struct {
	Prefix     string `json:"prefix"`
	ASN        string `json:"origin_asn"`
	State      string `json:"rpki_state"`
	Description string `json:"description"`
}

type ValidityData struct {
    State       string `json:"rpki_state"`
    Description string `json:"description"`
}

type PrefixResponse struct {
	Prefix string `json:"prefix"`
}

func RPKISearchController(c *fiber.Ctx) error {
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

	var results []RouteValidity

	prefixes, err := fetchPrefixes(asn)
	if err != nil {
		return handleError(c, "Error making API request to prefixes endpoint:", err)
	}

	for _, p := range prefixes {
		validity, err := fetchValidity(asn, p.Prefix)
		if err != nil {
			return handleError(c, "Error making API request:", err)
		}

		results = append(results, RouteValidity{
			Prefix:     p.Prefix,
			ASN:        asn,
			State:      validity.State,
			Description: validity.Description,
		})
	}

	return c.Status(200).JSON(results)
}

func fetchPrefixes(asn string) ([]PrefixResponse, error) {
	appConfig := config.AppConfigInstance

	prefixesResponse, _, err := makeAPIRequest(fmt.Sprintf("http://localhost%s/api/prefixes?q=%s", appConfig.HTTPPort, asn))
	if err != nil {
		return nil, err
	}

	var prefixes []PrefixResponse
	if err := json.Unmarshal(prefixesResponse, &prefixes); err != nil {
		return nil, err
	}

	return prefixes, nil
}

func fetchValidity(asn, prefix string) (ValidityData, error) {
	url := fmt.Sprintf("https://rpki-validator.ripe.net/api/v1/validity/%s/%s", asn, prefix)
	responseBody, _, err := makeAPIRequest(url)

	if err != nil {
		return ValidityData{}, err
	}

	var responseStruct struct {
		ValidatedRoute struct {
			Validity struct {
				State       string `json:"state"`
				Description string `json:"description"`
			} `json:"validity"`
		} `json:"validated_route"`
	}

	if err := json.Unmarshal(responseBody, &responseStruct); err != nil {
		return ValidityData{}, err
	}

	return ValidityData{
		State:       responseStruct.ValidatedRoute.Validity.State,
		Description: responseStruct.ValidatedRoute.Validity.Description,
	}, nil
}

func handleError(c *fiber.Ctx, message string, err error) error {
	return c.Status(500).JSON(map[string]interface{}{
		"message": fmt.Sprintf("%s %v", message, err),
		"debug": map[string]interface{}{
			"error": true,
			"code":  500,
		},
	})
}

