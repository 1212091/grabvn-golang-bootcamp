package models

type PassengerFeedback struct {
	BookingCode string `gorm:"unique"`
	PassengerID int32
	Feedback    string
}
