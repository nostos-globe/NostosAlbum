package service

import (
	"encoding/json"
	"fmt"
	"main/internal/models"
	"net/http"
)

type TripClient struct {
	BaseURL string
}

func (c *TripClient) GetTripByID(token string, TripMediaID uint) (models.TripMedia, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/trips/%d", c.BaseURL, TripMediaID), nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return models.TripMedia{}, err
	}

	req.Header.Set("Cookie", "auth_token="+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return models.TripMedia{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("failed to get TripMedia: %d", resp.StatusCode)
		fmt.Printf("Error: %v\n", err)
		return models.TripMedia{}, err
	}

	var response models.TripMedia
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
		return models.TripMedia{}, err
	}

	return response, nil
}
