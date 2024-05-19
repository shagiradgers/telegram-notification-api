.generate-pb:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        api/telegram_notification.proto

.run:
	go run cmd/api.go

generate: .generate-pb

run: .run