package model

type Cuisine int

const (
	SouthIndian Cuisine = iota + 1
	NorthIndian
	Chinese
	Indian
	Italian
	// add more cuisines here
)
