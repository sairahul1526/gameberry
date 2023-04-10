package util

import (
	MODEL "gameberry/model"
	"time"
)

func GetRestaurantRecommendations(user MODEL.User, restaurants []MODEL.Restaurant) []string {
	// user.Cuisines and user.CostBrackets are sorted by NoOfOrders in descending order and consists of at least 3 elements - first element is the primary cuisine/cost, second and third element is the secondary cuisine/cost

	// sorting orders
	orders := []struct {
		name      string
		condition func(MODEL.Restaurant, int) (count int, add bool)
	}{
		{"Featured restaurants of primary cuisine and primary cost bracket", func(r MODEL.Restaurant, previousCount int) (int, bool) {
			return -1, r.IsRecommended && r.Cuisine == user.Cuisines[0].Type && r.CostBracket == user.CostBrackets[0].Type
		}},
		{"If none, then all featured restaurants of primary cuisine, secondary cost and secondary cuisine, primary cost", func(r MODEL.Restaurant, previousCount int) (int, bool) {
			if previousCount > 0 {
				return -1, false
			}
			return -1, r.IsRecommended && ((r.Cuisine == user.Cuisines[0].Type && (r.CostBracket == user.CostBrackets[1].Type || r.CostBracket == user.CostBrackets[2].Type)) || ((r.Cuisine == user.Cuisines[1].Type || r.Cuisine == user.Cuisines[2].Type) && r.CostBracket == user.CostBrackets[0].Type))
		}},
		{"All restaurants of Primary cuisine, primary cost bracket with rating >= 4", func(r MODEL.Restaurant, previousCount int) (int, bool) {
			return -1, r.Cuisine == user.Cuisines[0].Type && r.CostBracket == user.CostBrackets[0].Type && r.Rating >= 4
		}},
		{"All restaurants of Primary cuisine, secondary cost bracket with rating >= 4.5", func(r MODEL.Restaurant, previousCount int) (int, bool) {
			return -1, r.Cuisine == user.Cuisines[0].Type && (r.CostBracket == user.CostBrackets[1].Type || r.CostBracket == user.CostBrackets[2].Type) && r.Rating >= 4.5
		}},
		{"All restaurants of secondary cuisine, primary cost bracket with rating >= 4.5", func(r MODEL.Restaurant, previousCount int) (int, bool) {
			return -1, (r.Cuisine == user.Cuisines[1].Type || r.Cuisine == user.Cuisines[2].Type) && r.CostBracket == user.CostBrackets[0].Type && r.Rating >= 4.5
		}},
		{"Top 4 newly created restaurants by rating", func(r MODEL.Restaurant, previousCount int) (int, bool) {
			return 4, time.Since(r.OnboardedTime) < 48*time.Hour
		}},
		{"All restaurants of Primary cuisine, primary cost bracket with rating < 4", func(r MODEL.Restaurant, previousCount int) (int, bool) {
			return -1, r.Cuisine == user.Cuisines[0].Type && r.CostBracket == user.CostBrackets[0].Type && r.Rating < 4
		}},
		{"All restaurants of Primary cuisine, secondary cost bracket with rating < 4.5", func(r MODEL.Restaurant, previousCount int) (int, bool) {
			return -1, r.Cuisine == user.Cuisines[0].Type && (r.CostBracket == user.CostBrackets[1].Type || r.CostBracket == user.CostBrackets[2].Type) && r.Rating < 4.5
		}},
		{"All restaurants of secondary cuisine, primary cost bracket with rating < 4.5", func(r MODEL.Restaurant, previousCount int) (int, bool) {
			return -1, (r.Cuisine == user.Cuisines[1].Type || r.Cuisine == user.Cuisines[2].Type) && r.CostBracket == user.CostBrackets[0].Type && r.Rating < 4.5
		}},
	}

	// Iterate through each sorting order and add the matching restaurants to the recommendations
	recommendations := []string{}
	recommendationsMap := map[string]struct{}{}
	previousCount := 0
	for _, order := range orders {
		matchingRestaurants := []MODEL.Restaurant{}
		for _, restaurant := range restaurants {
			if _, ok := recommendationsMap[restaurant.ID]; !ok {
				if count, add := order.condition(restaurant, previousCount); add {
					matchingRestaurants = append(matchingRestaurants, restaurant)
					if count > 0 && len(matchingRestaurants) >= count {
						break
					}
				}
			}
		}

		previousCount = 0
		// Add the matching restaurants to the recommendations
		for _, restaurant := range matchingRestaurants {
			if _, ok := recommendationsMap[restaurant.ID]; !ok {
				recommendationsMap[restaurant.ID] = struct{}{}
				recommendations = append(recommendations, restaurant.ID)
				previousCount++
			}
		}
	}

	// add rest all restaurants
	for _, restaurant := range restaurants {
		if _, ok := recommendationsMap[restaurant.ID]; !ok {
			recommendationsMap[restaurant.ID] = struct{}{}
			recommendations = append(recommendations, restaurant.ID)
		}
	}

	return recommendations
}
