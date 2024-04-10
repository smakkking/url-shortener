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


.PHONY: shutdown
shutdown:
	docker-compose down -v

.PHONY: test
test:
	go test -v -count=1 ./...

apply-migrations:
	docker build -t migrator ./db
	sleep 5
	docker run --network host migrator  \
	-path=/migrations/ \
	-database "postgresql://postgres:postgres@localhost:7557/urls?sslmode=disable" up

build-docker:
	docker-compose up --build
	 	
	