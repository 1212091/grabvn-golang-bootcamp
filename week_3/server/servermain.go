package main

import (
	"context"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pb "../passengerfeedback"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	port                 = ":15000"
)

type Passenger struct {
	passengerID int
}

var db *gorm.DB

type PassengerFeedback struct {
	BookingCode string `gorm:"unique"`
	PassengerID int32
	Feedback    string
}

// PassengerFeedbackServer is used to implement passenger feedback system
type PassengerFeedbackServer struct{}

// AddPassengerFeedback implements add passenger feedback
func (passengerFeedbackServer *PassengerFeedbackServer) AddPassengerFeedback(ctx context.Context,
	addPassengerFeedbackRequest *pb.AddPassengerFeedbackRequest) (*pb.AddPassengerFeedbackResponse, error) {

	passengerFeedback := PassengerFeedback{
		BookingCode: addPassengerFeedbackRequest.PassengerFeedback.BookingCode,
		PassengerID: addPassengerFeedbackRequest.PassengerFeedback.PassengerID,
		Feedback:    addPassengerFeedbackRequest.PassengerFeedback.Feedback,
	}

	err := db.Create(&passengerFeedback).Error
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &pb.AddPassengerFeedbackResponse{PassengerFeedback: &pb.PassengerFeedback{
		BookingCode: passengerFeedback.BookingCode,
		PassengerID: passengerFeedback.PassengerID,
		Feedback:    passengerFeedback.Feedback,
	}, Success: true}, nil
}

// GetPassengerFeedbackByPassengerId implements get passenger feedback by passenger Id
func (passengerFeedbackServer *PassengerFeedbackServer) GetPassengerFeedbackByPassengerId(ctx context.Context,
	getPassengerFeedbackByPassengerRequest *pb.GetPassengerFeedbackByPassengerRequest) (*pb.GetPassengerFeedbackByPassengerIdResponse, error) {

	passengerFeedbacks := make([]*PassengerFeedback, 0)

	err := db.Where("passenger_id = ?", getPassengerFeedbackByPassengerRequest.PassengerID).Find(&passengerFeedbacks).Error
	if err != nil {
		log.Fatal(err)
		return nil, status.Error(codes.Unknown, "Unknown error from database")
	}

	passengerFeedbackChosenList := make([]*pb.PassengerFeedback, 0)

	for _, passengerFeedback := range passengerFeedbacks {
		passengerFeedbackResponse := &pb.PassengerFeedback{
			BookingCode: passengerFeedback.BookingCode,
			PassengerID: passengerFeedback.PassengerID,
			Feedback:    passengerFeedback.Feedback,
		}
		passengerFeedbackChosenList = append(passengerFeedbackChosenList, passengerFeedbackResponse)
	}
	return &pb.GetPassengerFeedbackByPassengerIdResponse{PassengerFeedbacks: passengerFeedbackChosenList}, nil
}

// DeletePassengerFeedbackByPassengerId implements delete passenger feedback by passenger Id
func (passengerFeedbackServer *PassengerFeedbackServer) DeletePassengerFeedbackByPassengerId(ctx context.Context,
	deletePassengerFeedbackByPassengerIdRequest *pb.DeletePassengerFeedbackByPassengerIdRequest) (*pb.DeletePassengerFeedbackByPassengerIdResponse, error) {

	err := db.Where("passenger_id = ?", deletePassengerFeedbackByPassengerIdRequest.PassengerId).Delete(PassengerFeedback{}).Error

	if err != nil {
		log.Fatal(err)
		return &pb.DeletePassengerFeedbackByPassengerIdResponse{Success: false}, status.Error(codes.Unknown, "Unknown error from database")
	}

	return &pb.DeletePassengerFeedbackByPassengerIdResponse{Success: true}, nil
}

// GetPassengerFeedbackByBookingCode implements get passenger feedback by booking code
func (passengerFeedbackServer *PassengerFeedbackServer) GetPassengerFeedbackByBookingCode(ctx context.Context,
	getPassengerFeedbackByBookingCodeRequest *pb.GetPassengerFeedbackByBookingCodeRequest) (*pb.GetPassengerFeedbackByBookingCodeResponse, error) {

	passengerFeedback := &pb.PassengerFeedback{}
	err := db.Where("booking_code = ?", getPassengerFeedbackByBookingCodeRequest.BookingCode).Find(&passengerFeedback).Error

	if err != nil {
		log.Fatal(err)
		return nil, status.Error(codes.Unknown, "Unknown error from database")
	}

	return &pb.GetPassengerFeedbackByBookingCodeResponse{PassengerFeedback: &pb.PassengerFeedback{
		BookingCode: passengerFeedback.BookingCode,
		PassengerID: passengerFeedback.PassengerID,
		Feedback:    passengerFeedback.Feedback,
	}}, nil
}

func main() {
	var err error
	db, err = gorm.Open("mysql", "dotran:leonardo1994@tcp(127.0.0.1:3306)/passenger?charset=utf8&parseTime=true")

	if err != nil {
		log.Fatal("failed to connect db")
	}

	db.LogMode(true)

	db.Debug().DropTableIfExists(PassengerFeedback{})

	err = db.AutoMigrate(PassengerFeedback{}).Error

	if err != nil {
		log.Fatal("failed to migrate table passenger feedback")
	}

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPassengerFeedbackServiceServer(grpcServer, &PassengerFeedbackServer{})
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
