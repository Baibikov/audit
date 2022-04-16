package clickhouse

import (
	"context"

	"audit/internal/app/types"
)

func (r Repository) Create(ctx context.Context, log types.Log) error {
	batch, err := r.conn.PrepareBatch(ctx, "insert into log")
	if err != nil {
		return err
	}

	err = batch.AppendStruct(&log)
	if err != nil {
		return err
	}

	return batch.Send()
}