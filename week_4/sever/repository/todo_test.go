// +build integration persistence

package repository

import (
	"testing"
	"time"

	"github.com/go-pg/pg"
	"github.com/stretchr/testify/suite"
	"github.com/xuanit/testing/todo/pb"
)

type ToDoRepositorySuite struct {
	db *pg.DB
	suite.Suite
	todoRep ToDoImpl
}

func (s *ToDoRepositorySuite) SetupSuite() {
	// Connect to PostgresQL
	s.db = pg.Connect(&pg.Options{
		User:                  "postgres",
		Password:              "example",
		Database:              "todo",
		Addr:                  "localhost" + ":" + "5433",
		RetryStatementTimeout: true,
		MaxRetries:            4,
		MinRetryBackoff:       250 * time.Millisecond,
	})

	// Create Table
	s.db.CreateTable(&pb.Todo{}, nil)

	s.todoRep = ToDoImpl{DB: s.db}
}

func (s *ToDoRepositorySuite) TearDownSuite() {
	s.db.DropTable(&pb.Todo{}, nil)
	s.db.Close()
}

// Handle all insert - get - delete checks
func (s *ToDoRepositorySuite) TestInsert() {
	item := &pb.Todo{Id: "new_item", Title: "meeting"}
	err := s.todoRep.Insert(item)

	s.Nil(err)

	newTodo, err := s.todoRep.Get(item.Id)
	s.Nil(err)
	s.Equal(item, newTodo)

	err = s.todoRep.Delete(item.Id)
	s.Nil(err)
}

func (s *ToDoRepositorySuite) TestList() {
	item1 := &pb.Todo{Id: "item_1", Title: "A", Completed: true}
	item2 := &pb.Todo{Id: "item_2", Title: "B", Completed: true}
	err1 := s.todoRep.Insert(item1)
	err2 := s.todoRep.Insert(item2)

	s.Nil(err1)
	s.Nil(err2)

	toDoList, err := s.todoRep.List(5, false)

	expectedToDoList := []*pb.Todo{item1, item2}

	s.Nil(err)
	s.Equal(expectedToDoList, toDoList)

	err1 = s.todoRep.Delete(item1.Id)
	err2 = s.todoRep.Delete(item2.Id)

	s.Nil(err1)
	s.Nil(err2)
}

func (s *ToDoRepositorySuite) TestListOutOfLimit() {
	item1 := &pb.Todo{Id: "item_1", Title: "A", Completed: true}
	item2 := &pb.Todo{Id: "item_2", Title: "B", Completed: true}
	err1 := s.todoRep.Insert(item1)
	err2 := s.todoRep.Insert(item2)

	s.Nil(err1)
	s.Nil(err2)

	toDoList, err := s.todoRep.List(1, false)
	expectedToDoList := []*pb.Todo{item1}

	s.Nil(err)
	s.Equal(expectedToDoList, toDoList)

	err1 = s.todoRep.Delete(item1.Id)
	err2 = s.todoRep.Delete(item2.Id)

	s.Nil(err1)
	s.Nil(err2)
}

func (s *ToDoRepositorySuite) TestDeleteExceptionCase() {
	item1 := &pb.Todo{Id: "item_1", Title: "A", Completed: true}

	err := s.todoRep.Delete(item1.Id)
	s.NotNil(err)
}

func (s *ToDoRepositorySuite) TestCreateExceptionCase() {
	result, err := s.todoRep.Get("-1")
	s.NotNil(err)
	s.Nil(nil, result)
}

func TestToDoRepository(t *testing.T) {
	suite.Run(t, new(ToDoRepositorySuite))
}
