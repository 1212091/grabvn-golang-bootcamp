package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "../passengerfeedback"
)

const (
	port = ":15000"
)

// PassengerFeedbackServer is used to implement passenger feedback system
type PassengerFeedbackServer struct {
	passengerFeedbacks []*pb.PassengerFeedback
}

// AddPassengerFeedback implements add passenger feedback
func (passengerFeedbackServer *PassengerFeedbackServer) AddPassengerFeedback(ctx context.Context,
	in *pb.AddPassengerFeedbackRequest) (*pb.AddPassengerFeedbackResponse, error) {

	passengerFeedbackServer.passengerFeedbacks = append(passengerFeedbackServer.passengerFeedbacks, in.PassengerFeedback)

	return &pb.AddPassengerFeedbackResponse{PassengerFeedback: in.PassengerFeedback, Success: true}, nil
}

// DeletePassengerFeedbackByPassengerId implements add passenger feedback
func (passengerFeedbackServer *PassengerFeedbackServer) DeletePassengerFeedbackByPassengerId(ctx context.Context,
	in *pb.DeletePassengerFeedbackByPassengerIdRequest) (*pb.DeletePassengerFeedbackByPassengerIdResponse, error) {

	return &pb.DeletePassengerFeedbackByPassengerIdResponse{Success: true}, nil
}

// DeletePassengerFeedbackByPassengerId implements add passenger feedback
func (passengerFeedbackServer *PassengerFeedbackServer) GetPassengerFeedbackByBookingCode(ctx context.Context,
	in *pb.GetPassengerFeedbackByBookingCodeRequest) (*pb.GetPassengerFeedbackByBookingCodeResponse, error) {

	return &pb.GetPassengerFeedbackByBookingCodeResponse{PassengerFeedback: nil}, nil
}

// DeletePassengerFeedbackByPassengerId implements add passenger feedback
func (passengerFeedbackServer *PassengerFeedbackServer) GetPassengerFeedbackByPassengerId(ctx context.Context,
	in *pb.GetPassengerFeedbackByPassengerRequest) (*pb.GetPassengerFeedbackByPassengerIdResponse, error) {

	return &pb.GetPassengerFeedbackByPassengerIdResponse{PassengerFeedbacks: nil}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPassengerFeedbackServiceServer(grpcServer, &PassengerFeedbackServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
