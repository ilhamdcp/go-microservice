package main

import (
	"context"
)

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Create(context.Context) error {
	return nil
}
