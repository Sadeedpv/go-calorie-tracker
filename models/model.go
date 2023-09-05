package models

type Calories struct{
	// Add pointer to make it optional data
	ID *int `json:"id"`
	Food string `json:"food"`
	Calorie int `json:"calorie"`
}