package models

type UserRegistered struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Login  string `json:"login"`
}

type RentStarted struct {
	UserID   string `json:"user_id"`
	CarID    string `json:"car_id"`
	RentTime string `json:"rent_time"`
}

type RentFinished struct {
	UserID        string `json:"user_id"`
	CarID         string `json:"car_id"`
	Email         string `json:"email"`
	TotalRentTime string `json:"rent_time"`
}
