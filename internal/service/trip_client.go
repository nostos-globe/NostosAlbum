package service

import (
    "encoding/json"
    "fmt"
    "net/http"
	"main/internal/models"
)

type TripClient struct {
    BaseURL string
}


type TripResponse struct {
    Trip models.Trip `json:"trip"`
}

func (c *TripClient) GetTripByID(token string, tripID uint) (TripResponse, error) {
    req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/trips/%d", c.BaseURL, tripID), nil)
    if err != nil {
        return TripResponse{}, err
    }

    req.Header.Set("Cookie", "auth_token="+token)
    
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return TripResponse{}, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return TripResponse{}, fmt.Errorf("failed to get trip: %d", resp.StatusCode)
    }

    var response TripResponse
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        return TripResponse{}, err
    }

    return response, nil
}
