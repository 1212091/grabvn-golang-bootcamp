package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"

	pb "../passengerfeedback"
)

const (
	address = "localhost:15000"
)

var (
	client pb.PassengerFeedbackServiceClient
)

func main() {
	// Set up a connection to the server.
	connection, err := grpc.Dial(address, grpc.WithInsecure())
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
	router.Run(":8088")
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
		_, ok := status.FromError(err)
		if ok {
			g.String(500, "Unexpected error")

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
		_, ok := status.FromError(err)
		if ok {
			g.String(500, "Unexpected error")

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
		_, ok := status.FromError(err)
		if ok {
			g.String(500, "Unexpected error")
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
		_, ok := status.FromError(err)
		if ok {
			g.String(500, err.Error())
		}
	} else {
		g.JSON(200, feedback.PassengerFeedback)
	}
}
