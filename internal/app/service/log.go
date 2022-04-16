package service

import (
	"context"

	"github.com/pkg/errors"

	"audit/internal/app/repository"
	"audit/internal/app/types"
)

type logger struct {
	storage *repository.Storage
}

func newLogger(storage *repository.Storage ) *logger {
	return &logger{
		storage: storage,
	}
}

func (l *logger) Log(ctx context.Context, log types.Log) error {
	return errors.Wrap(l.storage.Log.Create(ctx, log), "audit logger, write to storage")
}