package model

type User struct {
	Cuisines     []CuisineTracking
	CostBrackets []CostTracking
}

type CuisineTracking struct {
	Type       Cuisine
	NoOfOrders int
}

type CostTracking struct {
	Type       int
	NoOfOrders int
}
