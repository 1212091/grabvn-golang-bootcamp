package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "../passengerfeedback"
)

const (
	SERVER_ADDRESS = "localhost:15000"
	REST_PORT = ":8088"
)

var (
	client pb.PassengerFeedbackServiceClient
)

func main() {
	// Set up a connection to the server.
	connection, err := grpc.Dial(SERVER_ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	defer connection.Close()

	client = pb.NewPassengerFeedbackServiceClient(connection)
	_, cancel := context.WithTimeout(context.Background(), 500*time.Second)
	defer cancel()

	router := gin.Default()
	router.PUT("/addPassengerFeedback", addPassengerFeedback)
	router.GET("/getPassengerFeedbackByPassengerId", getPassengerFeedbackByPassengerId)
	router.GET("/getPassengerFeedbackByBookingCode", getPassengerFeedbackByBookingCode)
	router.DELETE("/deletePassengerFeedbackByPassengerId", deletePassengerFeedbackByPassengerId)
	router.Run(REST_PORT)
}

func deletePassengerFeedbackByPassengerId(g *gin.Context) {
	var argument struct {
		PassengerID int32
	}

	err := g.BindJSON(&argument)

	if err != nil {
		g.String(400, "Bad request")
		return
	}

	deletePassengerFeedbackRequest := &pb.DeletePassengerFeedbackByPassengerIdRequest{
		PassengerId: argument.PassengerID,
	}

	response, err := client.DeletePassengerFeedbackByPassengerId(context.Background(), deletePassengerFeedbackRequest)
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			switch status.Code() {
			case codes.NotFound:
				g.JSON(200, "Not existed")
			default:
				g.JSON(500, "Unknown errror from database")
			}
		}
	} else {
		g.JSON(200, response)
	}
}

func getPassengerFeedbackByBookingCode(g *gin.Context) {
	bookingCode := g.Query("bookingCode")

	getPassengerFeedbackByBookingCodeRequest := &pb.GetPassengerFeedbackByBookingCodeRequest{
		BookingCode: bookingCode,
	}

	feedback, err := client.GetPassengerFeedbackByBookingCode(context.Background(), getPassengerFeedbackByBookingCodeRequest)
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			switch status.Code() {
			case codes.NotFound:
				feedback := pb.GetPassengerFeedbackByBookingCodeResponse{}
				g.JSON(200, feedback)
			default:
				g.JSON(500, "Unknown errror from database")
			}
		}
	} else {
		g.JSON(200, feedback)
	}
}

func getPassengerFeedbackByPassengerId(g *gin.Context) {
	param := g.Query("passengerId")
	passengerId, err := strconv.Atoi(param)

	if err != nil {
		g.String(400, "Bad request")
		return
	}

	getPassengerFeedbackByPassengerRequest := &pb.GetPassengerFeedbackByPassengerRequest{
		PassengerID: int32(passengerId),
	}

	feedbacks, err := client.GetPassengerFeedbackByPassengerId(context.Background(), getPassengerFeedbackByPassengerRequest)
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			switch status.Code() {
			case codes.NotFound:
				feedbacks = &pb.GetPassengerFeedbackByPassengerIdResponse{PassengerFeedbacks: make([]*pb.PassengerFeedback, 0)}
				g.JSON(200, feedbacks)
			default:
				g.JSON(500, "Unknown errror from database")
			}
		}
	} else {
		g.JSON(200, feedbacks)
	}

}

func addPassengerFeedback(g *gin.Context) {
	var argument struct {
		BookingCode string
		PassengerID int32
		Feedback    string
	}

	err := g.BindJSON(&argument)

	if err != nil {
		g.String(400, "Bad request")
		return
	}

	addPassengerFeedbackRequest := &pb.AddPassengerFeedbackRequest{
		PassengerFeedback: &pb.PassengerFeedback{
			BookingCode: argument.BookingCode,
			PassengerID: argument.PassengerID,
			Feedback:    argument.Feedback,
		},
	}

	feedback, err := client.AddPassengerFeedback(context.Background(), addPassengerFeedbackRequest)
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			switch status.Code() {
			case codes.AlreadyExists:
				g.JSON(500, "Feedback uses existing booking code")
			default:
				g.JSON(500, "Unknown errror from database")
			}
		}
	} else {
		g.JSON(200, feedback.PassengerFeedback)
	}
}
