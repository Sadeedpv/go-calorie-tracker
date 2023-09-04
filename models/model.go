package models

type Calories struct{
	ID int `json:"id"`
	Food string `json:"food"`
	Calorie int `json:"calorie"`
}