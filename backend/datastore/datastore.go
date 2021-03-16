package datastore

import (
	"errors"
	"sync"

	"github.com/rs/xid"
)

type StorerI interface {
	TodoList() []*Todo
	TodoCreate(*Todo) *Todo
	TodoDelete(id string) error
}

type Todo struct {
	Id   string `json:"id"`
	Body string `json:"body"`
}

func NewDatastore(todos []*Todo) StorerI {
	return &DataStore{data: todos}
}

type DataStore struct {
	sync.RWMutex
	data []*Todo
}

func (s *DataStore) TodoList() []*Todo {
	s.RLock()
	defer s.RUnlock()
	return s.data
}

func (s *DataStore) TodoCreate(todo *Todo) *Todo {
	if todo.Id == "" {
		todo.Id = xid.New().String()
	}
	s.Lock()
	defer s.Unlock()
	s.data = append(s.data, todo)
	return todo
}

func (s *DataStore) TodoDelete(id string) error {
	s.Lock()
	defer s.Unlock()
	for i, todo := range s.data {
		if todo.Id == id {
			s.data = append(s.data[:i], s.data[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
