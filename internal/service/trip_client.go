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


func (c *TripClient) GetTripByID(token string, tripID uint) (models.Trip, error) {
    fmt.Printf("Getting trip with ID: %d\n", tripID)
    
    req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/trips/%d", c.BaseURL, tripID), nil)
    if err != nil {
        fmt.Printf("Error creating request: %v\n", err)
        return models.Trip{}, err
    }

    req.Header.Set("Cookie", "auth_token="+token)
    fmt.Printf("Making request to: %s\n", req.URL.String())
    
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Printf("Error making request: %v\n", err)
        return models.Trip{}, err
    }
    defer resp.Body.Close()

    fmt.Printf("Response status code: %d\n", resp.StatusCode)
    if resp.StatusCode != http.StatusOK {
        err := fmt.Errorf("failed to get trip: %d", resp.StatusCode)
        fmt.Printf("Error: %v\n", err)
        return models.Trip{}, err
    }

    var response models.Trip
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        fmt.Printf("Error decoding response: %v\n", err)
        return models.Trip{}, err
    }
    fmt.Printf("Response: %+v\n", response)
    fmt.Printf("Successfully retrieved trip with ID: %d\n", response.TripID)
    return response, nil
}
