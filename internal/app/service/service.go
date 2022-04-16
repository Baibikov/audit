package service

import (
	"context"

	"audit/internal/app/repository"
	"audit/internal/app/types"
)

type UseCase struct {
	Logger Logger
}

type Logger interface {
	Log(ctx context.Context, log types.Log) error
}

func New(storage *repository.Storage) *UseCase {
	return &UseCase{
		Logger: newLogger(storage),
	}
}