package model

// FoodDB db data
type FoodDB struct {
	Good []Food `json:"good"`
	Bad  []Food `json:"bad"`
}

// Food food
type Food struct {
	Name  string `json:"name"`
	Where string `json:"where"`
	Why   string `json:"why"`
}
