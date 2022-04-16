proto:
	mkdir -p ./internal/app/delivery/grpc/generated
	protoc  --go_out=. --go-grpc_out=. ./internal/app/delivery/grpc/audit_logger_service.proto