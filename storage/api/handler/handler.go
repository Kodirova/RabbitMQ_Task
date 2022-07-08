package handler

import (
	"test_task/storage/internal"
)

type handler struct {
	storage internal.Storage
}

type HandlerOptions struct {
	Storage internal.Storage
}

func New(options *HandlerOptions) *handler {
	return &handler{
		storage: options.Storage,
	}
}
