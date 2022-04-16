package repository

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"

	"audit/internal/app/repository/clickhouse"
	"audit/internal/app/types"
)

type Storage struct {
	Log Log
}

type Log interface {
	Create(ctx context.Context, log types.Log) error
}

func New(conn driver.Conn) *Storage{
	return &Storage{
		Log: clickhouse.New(conn),
	}
}
