package util

import (
	MODEL "gameberry/model"
	"reflect"
	"testing"
	"time"
)

func TestGetRestaurantRecommendations(t *testing.T) {
	// Example user
	user := MODEL.User{
		Cuisines: []MODEL.CuisineTracking{
			{Type: MODEL.Indian, NoOfOrders: 10},
			{Type: MODEL.Chinese, NoOfOrders: 5},
			{Type: MODEL.Italian, NoOfOrders: 3},
		},
		CostBrackets: []MODEL.CostTracking{
			{Type: 1, NoOfOrders: 15},
			{Type: 2, NoOfOrders: 8},
			{Type: 3, NoOfOrders: 5},
		},
	}

	// Example restaurants
	restaurants := []MODEL.Restaurant{
		{
			ID:            "1",
			Cuisine:       MODEL.Indian,
			CostBracket:   1,
			Rating:        4.5,
			IsRecommended: true,
			OnboardedTime: time.Now().Add(-24 * time.Hour),
		},
		{
			ID:            "2",
			Cuisine:       MODEL.Chinese,
			CostBracket:   2,
			Rating:        4.6,
			IsRecommended: false,
			OnboardedTime: time.Now().Add(-24 * time.Hour),
		},
		{
			ID:            "3",
			Cuisine:       MODEL.Italian,
			CostBracket:   3,
			Rating:        3.5,
			IsRecommended: true,
			OnboardedTime: time.Now().Add(-24 * time.Hour),
		},
		{
			ID:            "4",
			Cuisine:       MODEL.Indian,
			CostBracket:   2,
			Rating:        4.0,
			IsRecommended: true,
			OnboardedTime: time.Now().Add(-48 * time.Hour),
		},
		{
			ID:            "5",
			Cuisine:       MODEL.Chinese,
			CostBracket:   3,
			Rating:        3.5,
			IsRecommended: false,
			OnboardedTime: time.Now().Add(-48 * time.Hour),
		},
		{
			ID:            "6",
			Cuisine:       MODEL.Indian,
			CostBracket:   3,
			Rating:        4.8,
			IsRecommended: true,
			OnboardedTime: time.Now().Add(-12 * time.Hour),
		},
	}

	// Expected output
	expected := []string{"1", "6", "2", "3", "4", "5"}

	// Test the function
	output := GetRestaurantRecommendations(user, restaurants)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("GetRestaurantRecommendations failed: expected %v, but got %v", expected, output)
	}
}
