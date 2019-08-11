// +build integration

package consumer

import (
	"log"
	"testing"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"

	"github.com/pact-foundation/pact-go/dsl"
)

func TestCreateToDo(t *testing.T) {
	// Create Pact connecting to local Daemon
	pact := &dsl.Pact{
		Consumer: "ToDoConsumer",
		Provider: "ToDoService",
		Host:     "localhost",
	}
	defer pact.Teardown()

	createToDoRes := struct {
		id string `json:"id" pact:"example=id1"`
	}{}
	// Set up our expected interactions.
	pact.
		AddInteraction().
		//Given("UserA is existing").
		UponReceiving("A request to create todo").
		WithRequest(dsl.Request{
			Method:  "POST",
			Path:    dsl.String("/v1/todo"),
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body: map[string]interface{}{
				"title":       "1-1 with manager",
				"description": "discuss about OKRs",
				"completed":   true,
			},
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body:    dsl.Match(createToDoRes),
		})

	// Pass in test case. This is the component that makes the external HTTP call
	var test = func() (err error) {
		proxy := ToDoProxy{Host: "localhost", Port: pact.Server.Port}
		id, err := proxy.CreateToDo(ToDo{Id: "1", Title: "1-1 with manager", Description: "discuss about OKRs", Completed: true})
		if err != nil {
			return err
		}
		assert.Equal(t, "id1", id)
		return nil
	}

	// Run the test, verify it did what we expected and capture the contract
	if err := pact.Verify(test); err != nil {
		log.Fatalf("Error on Verify: %v", err)
	}
}

func TestGetToDo(t *testing.T) {
	// Create Pact connecting to local Daemon
	pact := &dsl.Pact{
		Consumer: "ToDoConsumer",
		Provider: "ToDoService",
		Host:     "localhost",
	}
	defer pact.Teardown()

	getToDoRes := struct {
		id          string               `json:"id" pact:"example=1"`
		Title       string               `json:"title" pact:"example=Job fair"`
		Description string               `json:"description" pact:"example=Description"`
		Completed   bool                 `json:"completed" pact:true`
		CreatedAt   *timestamp.Timestamp `json:"created_at,omitempty"`
		UpdatedAt   *timestamp.Timestamp `json:"updated_at,omitempty"`
	}{}
	// Set up our expected interactions.
	pact.
		AddInteraction().
		UponReceiving("A request to get todo").
		WithRequest(dsl.Request{
			Method:  "GET",
			Path:    dsl.String("/v1/todo/1"),
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body:    dsl.Match(&getToDoRes),
		})

	// Pass in test case. This is the component that makes the external HTTP call

	var test = func() (err error) {
		proxy := ToDoProxy{Host: "localhost", Port: pact.Server.Port}
		res, err := proxy.GetToDo("1")
		if err != nil {
			return err
		}
		expectedRes := ToDo{
			Id:          "1",
			Title:       "Job fair",
			Description: "Description",
			Completed:   true,
			CreatedAt:   nil,
			UpdatedAt:   nil,
		}
		assert.Equal(t, expectedRes, res)
		return nil
	}

	// Run the test, verify it did what we expected and capture the contract
	if err := pact.Verify(test); err != nil {
		log.Fatalf("Error on Verify: %v", err)
	}
}
