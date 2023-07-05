package db

import (
	"toleg/internal/client/admin"
	"toleg/pkg/client/postgresql"
	"toleg/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) admin.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
