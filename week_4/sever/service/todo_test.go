package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuanit/testing/todo/pb"
	"github.com/xuanit/testing/todo/server/repository/mocks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func TestGetToDo(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}
	toDo := &pb.Todo{}
	req := &pb.GetTodoRequest{Id: "123"}
	mockToDoRep.On("Get", req.Id).Return(toDo, nil)
	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.GetTodo(nil, req)

	expectedRes := &pb.GetTodoResponse{Item: toDo}

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockToDoRep.AssertExpectations(t)
}

func TestGetToDoExceptionCase(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}
	toDo := &pb.Todo{}
	req := &pb.GetTodoRequest{Id: "123"}
	mockErr := errors.New("")
	mockToDoRep.On("Get", req.Id).Return(toDo, mockErr)
	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.GetTodo(nil, req)

	var expectedRes *pb.GetTodoResponse = nil
	expectedErr := grpc.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", mockErr)

	assert.NotNil(t, err)
	assert.Equal(t, expectedRes, res)
	assert.Equal(t, expectedErr, err)
	mockToDoRep.AssertExpectations(t)
}

func TestListToDo(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}

	mockToDoList := []*pb.Todo{
		&pb.Todo{
			Id:          "123",
			Title:       "abc",
			Description: "xyz",
			Completed:   true,
			CreatedAt:   nil,
		},
	}

	req := &pb.ListTodoRequest{Limit: 10, NotCompleted: false}
	mockToDoRep.On("List", req.Limit, req.NotCompleted).Return(mockToDoList, nil)
	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.ListTodo(nil, req)

	expectedRes := &pb.ListTodoResponse{Items: []*pb.Todo{
		&pb.Todo{
			Id:          "123",
			Title:       "abc",
			Description: "xyz",
			Completed:   true,
			CreatedAt:   nil,
		},
	}}

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockToDoRep.AssertExpectations(t)
}

func TestListToDoExceptionCase(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}
	mockToDoList := []*pb.Todo{}
	req := &pb.ListTodoRequest{Limit: 10, NotCompleted: false}
	mockErr := errors.New("")
	mockToDoRep.On("List", req.Limit, req.NotCompleted).Return(mockToDoList, mockErr)
	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.ListTodo(nil, req)

	var expectedRes *pb.ListTodoResponse = nil
	expectedErr := grpc.Errorf(codes.NotFound, "Could not list items from the database: %s", mockErr)

	assert.NotNil(t, err)
	assert.Equal(t, expectedRes, res)
	assert.Equal(t, expectedErr, err)
	mockToDoRep.AssertExpectations(t)
}

func TestCreateToDo(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}
	mockToDo := &pb.Todo{
		Id:          "123",
		Title:       "abc",
		Description: "xyz",
		Completed:   true,
		CreatedAt:   nil,
	}
	req := &pb.CreateTodoRequest{Item: mockToDo}
	mockToDoRep.On("Insert", req.Item).Return(nil)
	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.CreateTodo(nil, req)

	expectedRes := &pb.CreateTodoResponse{
		Id: mockToDo.Id,
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockToDoRep.AssertExpectations(t)
}

func TestCreateToDoExceptionCase(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}
	mockToDo := &pb.Todo{}
	req := &pb.CreateTodoRequest{Item: mockToDo}
	mockErr := errors.New("")
	mockToDoRep.On("Insert", req.Item).Return(mockErr)
	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.CreateTodo(nil, req)

	expectedErr := grpc.Errorf(codes.Internal, "Could not insert item into the database: %s", mockErr)
	var expectedRes *pb.CreateTodoResponse = nil

	assert.NotNil(t, err)
	assert.Equal(t, expectedRes, res)
	assert.Equal(t, expectedErr, err)
	mockToDoRep.AssertExpectations(t)
}

func TestDeleteToDo(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}
	req := &pb.DeleteTodoRequest{Id: "a1"}
	mockToDoRep.On("Delete", req.Id).Return(nil)
	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.DeleteTodo(nil, req)

	expectedRes := &pb.DeleteTodoResponse{}
	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockToDoRep.AssertExpectations(t)
}

func TestDeleteToDoExceptionCase(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}
	req := &pb.DeleteTodoRequest{Id: "a2"}
	mockErr := errors.New("")
	mockToDoRep.On("Delete", req.Id).Return(mockErr)
	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.DeleteTodo(nil, req)

	expectedErr := grpc.Errorf(codes.Internal, "Could not delete item from the database: %s", mockErr)
	var expectedRes *pb.DeleteTodoResponse = nil

	assert.NotNil(t, err)
	assert.Equal(t, expectedRes, res)
	assert.Equal(t, expectedErr, err)
	mockToDoRep.AssertExpectations(t)
}
