package main

import (
	"context"
	"errors"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "../passengerfeedback"
)

const (
	port = ":15000"
)

var passengerList = []Passenger{Passenger{passengerID: 1}, Passenger{passengerID: 2}, Passenger{passengerID: 3}, Passenger{passengerID: 4}}

var passengerFeedbackList = make([]PassengerFeedback, 0)

type Passenger struct {
	passengerID int
}

// PassengerFeedbackServer is used to implement passenger feedback system
type PassengerFeedback struct {
	passengerFeedback *pb.PassengerFeedback
}

// AddPassengerFeedback implements add passenger feedback
func (passengerFeedbackServer *PassengerFeedback) AddPassengerFeedback(ctx context.Context,
	addPassengerFeedbackRequest *pb.AddPassengerFeedbackRequest) (*pb.AddPassengerFeedbackResponse, error) {

	err := validateBookingCode(addPassengerFeedbackRequest)

	if err != nil {
		return &pb.AddPassengerFeedbackResponse{PassengerFeedback: passengerFeedbackList[len(passengerFeedbackList)-1].passengerFeedback, Success: false}, err
	}

	passengerFeedbackList = append(passengerFeedbackList, PassengerFeedback{passengerFeedback: addPassengerFeedbackRequest.PassengerFeedback})

	return &pb.AddPassengerFeedbackResponse{PassengerFeedback: passengerFeedbackList[len(passengerFeedbackList)-1].passengerFeedback, Success: true}, nil
}

// DeletePassengerFeedbackByPassengerId implements delete passenger feedback by passenger Id
func (passengerFeedbackServer *PassengerFeedback) DeletePassengerFeedbackByPassengerId(ctx context.Context,
	deletePassengerFeedbackByPassengerIdRequest *pb.DeletePassengerFeedbackByPassengerIdRequest) (*pb.DeletePassengerFeedbackByPassengerIdResponse, error) {

	for index, passengerFeedback := range passengerFeedbackList {
		if passengerFeedback.passengerFeedback.PassengerID == deletePassengerFeedbackByPassengerIdRequest.PassengerId {
			removePassengerFeedbackOutOfList(index)
		}
	}

	return &pb.DeletePassengerFeedbackByPassengerIdResponse{Success: true}, nil
}

func removePassengerFeedbackOutOfList(index int) {
	passengerFeedbackList = append(passengerFeedbackList[:index], passengerFeedbackList[index+1:]...)
}

func validateBookingCode(addPassengerFeedbackRequest *pb.AddPassengerFeedbackRequest) error {
	for _, passengerFeedback := range passengerFeedbackList {
		if passengerFeedback.passengerFeedback.BookingCode == addPassengerFeedbackRequest.PassengerFeedback.BookingCode {
			err := errors.New("The booking code of this feedback is already existed")
			return err
		}
	}
	return nil
}

// GetPassengerFeedbackByBookingCode implements get passenger feedback by booking code
func (passengerFeedbackServer *PassengerFeedback) GetPassengerFeedbackByBookingCode(ctx context.Context,
	getPassengerFeedbackByBookingCodeRequest *pb.GetPassengerFeedbackByBookingCodeRequest) (*pb.GetPassengerFeedbackByBookingCodeResponse, error) {

	for _, passengerFeedback := range passengerFeedbackList {
		if passengerFeedback.passengerFeedback.BookingCode == getPassengerFeedbackByBookingCodeRequest.BookingCode {
			return &pb.GetPassengerFeedbackByBookingCodeResponse{PassengerFeedback: passengerFeedback.passengerFeedback}, nil
		}
	}

	return &pb.GetPassengerFeedbackByBookingCodeResponse{PassengerFeedback: nil}, nil
}

// GetPassengerFeedbackByPassengerId implements get passenger feedback by passenger Id
func (passengerFeedbackServer *PassengerFeedback) GetPassengerFeedbackByPassengerId(ctx context.Context,
	getPassengerFeedbackByPassengerRequest *pb.GetPassengerFeedbackByPassengerRequest) (*pb.GetPassengerFeedbackByPassengerIdResponse, error) {

	passengerFeedbackChosenList := make([]*pb.PassengerFeedback, 0)

	for _, passengerFeedback := range passengerFeedbackList {
		if passengerFeedback.passengerFeedback.PassengerID == getPassengerFeedbackByPassengerRequest.PassengerID {
			passengerFeedbackChosenList = append(passengerFeedbackChosenList, passengerFeedback.passengerFeedback)
		}
	}

	return &pb.GetPassengerFeedbackByPassengerIdResponse{PassengerFeedbacks: passengerFeedbackChosenList}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPassengerFeedbackServiceServer(grpcServer, &PassengerFeedback{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
