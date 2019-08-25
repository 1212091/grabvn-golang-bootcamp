package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

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
	fmt.Println("Choose 1 if you want to apply retry in circuit breaker - 2 if you want to apply circuit breaker in retry: ")
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		choice, err := strconv.ParseInt(input, 10, 64)
		if err == nil {
			if choice == 1 {
				client.beginRequestWithRetryInCircuitBreaker()
			} else if choice == 2 {
				client.beginRequestWithCircuitBreakerInRetry()
			}
		}
	}
}

func (client *Client) beginRequestWithCircuitBreakerInRetry() {
	for i := 1; i <= numberOfSampleRequest; i++ {
		j := i
		retry.Do(
			func() error {
				client.logger.Info("Request ", j, " begin")
				err := client.callWithCircuitBreaker(j)
				return err
			},
			retry.Attempts(maxAttempts),
		)
	}
}

func (client *Client) callWithCircuitBreaker(j int) error {
	err := hystrix.Do(circuitBreakerName, func() error {
		resp, err := http.Get(url)
		if resp.StatusCode >= 500 {
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
	}, func(err error) error {
		client.logger.Error("request ", j, ": fail, service is not available\n")
		return errors.New("Fall back")
	})
	return err
}

func (client *Client) beginRequestWithRetryInCircuitBreaker() {
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
			if resp.StatusCode >= 500 {
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
