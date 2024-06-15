package main

import "context"

func NewService(store OrderStore) *Service {
	return &Service{store: store}
}

func (s *Service) CreateOrder(context.Context) error {
	return nil
}
