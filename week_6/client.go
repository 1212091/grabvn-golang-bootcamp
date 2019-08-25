package main

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/avast/retry-go"
	hystrix "github.com/myteksi/hystrix-go/hystrix"
	log "github.com/sirupsen/logrus"
)

const (
	url                   = "http://localhost:15000/test"
	maxAttempts           = 2
	numberOfSampleRequest = 100
	circuitBreakerName    = "my_circuit_breaker"
)

type Client struct {
	logger *log.Logger
}

func main() {
	logger := log.New()
	client := Client{
		logger: logger,
	}
	client.initLogLevel()
	client.configureCircuitBreaker()
	client.beginRequest()

}

func (client *Client) beginRequest() {
	for i := 1; i <= numberOfSampleRequest; i++ {
		j := i
		hystrix.Do(circuitBreakerName, func() error {
			client.logger.Info("Request ", j, " begin")
			err := client.callWithRetry()
			return err
		}, func(err error) error {
			client.logger.Error("request ", j, ": fail, service is not available\n")
			return errors.New("Fall back")
		})
	}
}

func (client *Client) initLogLevel() {
	client.logger.SetLevel(log.InfoLevel)
}

func (client *Client) callWithRetry() error {
	err := retry.Do(
		func() error {
			resp, err := http.Get(url)
			if resp.StatusCode != 200 {
				client.logger.Error("Server error")
				return errors.New("Server error")
			}
			defer resp.Body.Close()
			_, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				client.logger.Error("Cannot parse response body")
				return errors.New("Cannot parse response body")
			}
			return nil
		},
		retry.Attempts(maxAttempts),
	)
	return err
}

func (client *Client) configureCircuitBreaker() {
	hystrix.ConfigureCommand(circuitBreakerName, hystrix.CommandConfig{
		ErrorPercentThreshold: 10,
	})
}
