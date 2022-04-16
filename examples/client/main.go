package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"audit/internal/app/delivery/grpc/generated"
)

const host = "127.0.0.1:8456"

var ctx = context.Background()

func main() {
	conn, err := grpc.Dial(
		host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln(err)
	}

	info := &generated.Info{
		ObjectCode: "test_object_code",
		ActionCode: "create_action",
		Data: map[string]string{
			"key": "value",
		},
		CreatedAt: time.Now().Unix(),
	}

	if _, err :=  generated.NewAuditLoggerClient(conn).Log(ctx, info); err != nil {
		log.Fatalln(err)
	}
}
