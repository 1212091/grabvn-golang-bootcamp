package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	feedbackrepository "./repositories"

	pb "../passengerfeedback"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	port = ":15000"
)

// PassengerFeedbackServer is used to implement passenger feedback system
type PassengerFeedbackServer struct{}

// AddPassengerFeedback implements add passenger feedback
func (passengerFeedbackServer *PassengerFeedbackServer) AddPassengerFeedback(ctx context.Context,
	addPassengerFeedbackRequest *pb.AddPassengerFeedbackRequest) (*pb.AddPassengerFeedbackResponse, error) {
	passengerFeedback, err := feedbackrepository.AddPassengerFeedback(addPassengerFeedbackRequest)

	return &pb.AddPassengerFeedbackResponse{PassengerFeedback: &pb.PassengerFeedback{
		BookingCode: passengerFeedback.BookingCode,
		PassengerID: passengerFeedback.PassengerID,
		Feedback:    passengerFeedback.Feedback,
	}, Success: true}, err
}

// GetPassengerFeedbackByPassengerId implements get passenger feedback by passenger Id
func (passengerFeedbackServer *PassengerFeedbackServer) GetPassengerFeedbackByPassengerId(ctx context.Context,
	getPassengerFeedbackByPassengerRequest *pb.GetPassengerFeedbackByPassengerRequest) (*pb.GetPassengerFeedbackByPassengerIdResponse, error) {

	passengerFeedbacks, err := feedbackrepository.GetPassengerFeedbackByPassengerId(getPassengerFeedbackByPassengerRequest)

	passengerFeedbackChosenList := make([]*pb.PassengerFeedback, 0)

	for _, passengerFeedback := range passengerFeedbacks {
		passengerFeedbackResponse := &pb.PassengerFeedback{
			BookingCode: passengerFeedback.BookingCode,
			PassengerID: passengerFeedback.PassengerID,
			Feedback:    passengerFeedback.Feedback,
		}
		passengerFeedbackChosenList = append(passengerFeedbackChosenList, passengerFeedbackResponse)
	}
	return &pb.GetPassengerFeedbackByPassengerIdResponse{PassengerFeedbacks: passengerFeedbackChosenList}, err
}

// DeletePassengerFeedbackByPassengerId implements delete passenger feedback by passenger Id
func (passengerFeedbackServer *PassengerFeedbackServer) DeletePassengerFeedbackByPassengerId(ctx context.Context,
	deletePassengerFeedbackByPassengerIdRequest *pb.DeletePassengerFeedbackByPassengerIdRequest) (*pb.DeletePassengerFeedbackByPassengerIdResponse, error) {
	err := feedbackrepository.DeletePassengerFeedbackByPassengerId(deletePassengerFeedbackByPassengerIdRequest)
	if (err != nil) {
		return &pb.DeletePassengerFeedbackByPassengerIdResponse{Success: false}, err
	} else {
		return &pb.DeletePassengerFeedbackByPassengerIdResponse{Success: true}, err
	}
}

// GetPassengerFeedbackByBookingCode implements get passenger feedback by booking code
func (passengerFeedbackServer *PassengerFeedbackServer) GetPassengerFeedbackByBookingCode(ctx context.Context,
	getPassengerFeedbackByBookingCodeRequest *pb.GetPassengerFeedbackByBookingCodeRequest) (*pb.GetPassengerFeedbackByBookingCodeResponse, error) {

	passengerFeedback, err := feedbackrepository.GetPassengerFeedbackByBookingCode(getPassengerFeedbackByBookingCodeRequest)

	return &pb.GetPassengerFeedbackByBookingCodeResponse{PassengerFeedback: &pb.PassengerFeedback{
		BookingCode: passengerFeedback.BookingCode,
		PassengerID: passengerFeedback.PassengerID,
		Feedback:    passengerFeedback.Feedback,
	}}, err
}

func main() {
	feedbackrepository.Init()

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
