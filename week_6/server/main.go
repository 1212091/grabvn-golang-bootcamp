package main

import (
	"math/rand"
	"net/http"

	statsd "github.com/DataDog/datadog-go/statsd"
	log "github.com/sirupsen/logrus"
)

var client = newClient()

func main() {

	// Create our server
	logger := log.New()
	server := Server{
		logger: logger,
	}
	newClient()
	// Start the server
	server.ListenAndServe()
}

func newClient() *statsd.Client {
	client, err := statsd.New("127.0.0.1:8125",
		statsd.WithNamespace("test."),                  // prefix every metric with the app name
		statsd.WithTags([]string{"region:us-east-1a"}), // send the EC2 availability zone as a tag with every metric
	)

	if err != nil {
		log.Fatal(err)
	}
	return client
}

// Server represents our server.
type Server struct {
	logger *log.Logger
}

// ListenAndServe starts the server
func (s *Server) ListenAndServe() {
	s.logger.Info("echo server is starting on port 15000...")
	http.HandleFunc("/", s.echo)
	http.ListenAndServe(":15000", nil)
}

// Echo echos back the request as a response
func (s *Server) echo(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")

	// 30% chance of failure
	if rand.Intn(100) < 30 {
		client.Count("fail", 1, []string{"localhost"}, 1)
		writer.WriteHeader(500)
		writer.Write([]byte("a chaos monkey broke your server"))
		return
	}
	client.Count("success", 1, []string{"localhost"}, 1)
	// Happy path
	writer.WriteHeader(200)
	request.Write(writer)
}
