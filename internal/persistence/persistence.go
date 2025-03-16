package persistence

import "errors"

type Store[T any] interface {
	Get(id string) (T, error)
	Set(id string, o T) error
}

type InMemStore[T any] struct {
	data map[string]T
}

func NewInMemStore[T any]() *InMemStore[T] {
	return &InMemStore[T]{
		data: make(map[string]T),
	}
}

func (s *InMemStore[T]) Get(id string) (T, error) {
	d, ok := s.data[id]
	if ok == false {
		return d, errors.New("Doesn't exist")
	}
	return d, nil
}

func (s *InMemStore[T]) Set(id string, o T) error {
	s.data[id] = o

	return nil
}
