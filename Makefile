APP_NAME?=my-app

# сборка отдельного приложения
clean:
	rm -f ${APP_NAME}

build: clean
	go build -mod vendor -o ${APP_NAME} ./cmd/service/service.go

run: build
	./${APP_NAME}

.PHONY: gen-grpc
gen-grpc:
	protoc -I ./pkg/sdk/proto ./pkg/sdk/proto/*.proto \
	--go_out=./pkg/sdk/go --go_opt=paths=source_relative \
	--go-grpc_out=./pkg --go_opt=paths=source_relative