package feedbackrepository

import (
	"log"

	pb "../../passengerfeedback"

	feedbackmodel "../models"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	UNIQUE_VIOLATE_RULE_ERROR_CODE = 1062
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open("mysql", "dotran:leonardo1994@tcp(127.0.0.1:3306)/passenger?charset=utf8&parseTime=true")

	if err != nil {
		log.Fatal("failed to connect db")
	}

	db.LogMode(true)

	db.Debug().DropTableIfExists(feedbackmodel.PassengerFeedback{})

	err = db.AutoMigrate(feedbackmodel.PassengerFeedback{}).Error

	if err != nil {
		log.Fatal("failed to migrate table passenger feedback")
	}

}

func AddPassengerFeedback(addPassengerFeedbackRequest *pb.AddPassengerFeedbackRequest) (feedbackmodel.PassengerFeedback, error) {
	passengerFeedback := feedbackmodel.PassengerFeedback{
		BookingCode: addPassengerFeedbackRequest.PassengerFeedback.BookingCode,
		PassengerID: addPassengerFeedbackRequest.PassengerFeedback.PassengerID,
		Feedback:    addPassengerFeedbackRequest.PassengerFeedback.Feedback,
	}

	err := db.Create(&passengerFeedback).Error
	if err != nil {
		log.Println(err)
		errorType, ok := err.(*mysql.MySQLError)
		if !ok {
			return feedbackmodel.PassengerFeedback{}, status.Error(codes.Unknown, err.Error())
		} else if errorType.Number == UNIQUE_VIOLATE_RULE_ERROR_CODE {
			return feedbackmodel.PassengerFeedback{}, status.Error(codes.AlreadyExists, err.Error())
		} else {
			return feedbackmodel.PassengerFeedback{}, status.Error(codes.Unknown, err.Error())
		}
	}

	return passengerFeedback, nil
}

func GetPassengerFeedbackByPassengerId(getPassengerFeedbackByPassengerRequest *pb.GetPassengerFeedbackByPassengerRequest) ([]*feedbackmodel.PassengerFeedback, error) {
	passengerFeedbacks := make([]*feedbackmodel.PassengerFeedback, 0)

	err := db.Where("passenger_id = ?", getPassengerFeedbackByPassengerRequest.PassengerID).Find(&passengerFeedbacks).Error
	if err != nil {
		log.Println(err)
		if gorm.IsRecordNotFoundError(err) {
			return make([]*feedbackmodel.PassengerFeedback, 0), status.Error(codes.NotFound, "Record not found")
		} else {
			return make([]*feedbackmodel.PassengerFeedback, 0), status.Error(codes.Unknown, "Unknown error from database")
		}
	}
	return passengerFeedbacks, nil
}

func DeletePassengerFeedbackByPassengerId(deletePassengerFeedbackByPassengerIdRequest *pb.DeletePassengerFeedbackByPassengerIdRequest) error {
	err := db.Where("passenger_id = ?", deletePassengerFeedbackByPassengerIdRequest.PassengerId).Delete(feedbackmodel.PassengerFeedback{}).Error

	if err != nil {
		log.Println(err)
		return status.Error(codes.Unknown, "Unknown error from database")
	}

	return nil
}

func GetPassengerFeedbackByBookingCode(getPassengerFeedbackByBookingCodeRequest *pb.GetPassengerFeedbackByBookingCodeRequest) (feedbackmodel.PassengerFeedback, error) {
	passengerFeedback := &feedbackmodel.PassengerFeedback{}
	err := db.Where("booking_code = ?", getPassengerFeedbackByBookingCodeRequest.BookingCode).Find(&passengerFeedback).Error

	if err != nil {
		log.Println(err)
		if gorm.IsRecordNotFoundError(err) {
			return feedbackmodel.PassengerFeedback{}, status.Error(codes.NotFound, "Record not found")
		} else {
			return feedbackmodel.PassengerFeedback{}, status.Error(codes.Unknown, "Unknown error from database")
		}
	}

	return *passengerFeedback, nil
}
