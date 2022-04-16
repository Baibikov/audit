package grpc

import (
	"context"
	"time"

	"audit/internal/app/delivery/grpc/generated"
	"audit/internal/app/service"
	"audit/internal/app/types"
)

type Instance struct {
	generated.AuditLoggerServer
	usecase *service.UseCase
}

func New(usecase *service.UseCase) *Instance {
	return &Instance{
		usecase: usecase,
	}
}

func (h Instance) Log(ctx context.Context, info *generated.Info) (*generated.Empty, error) {
	err := h.usecase.Logger.Log(ctx, types.Log{
		ObjectCode: info.ObjectCode,
		ActionCode: info.ActionCode,
		Data: info.Data,
		CreatedAt: time.Unix(info.CreatedAt, 0),
	})
	return &generated.Empty{}, err
}