package controller

import (
	"github.com/l-orlov/simple-todo-list/server/store"
)

type Controller struct {
	storage *store.Storage
}

func New(storage *store.Storage) (*Controller, error) {
	return &Controller{storage: storage}, nil
}
