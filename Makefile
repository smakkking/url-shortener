APP_NAME?=my-app

# сборка отдельного приложения
clean:
	rm -f ${APP_NAME}

build: clean
	go build -mod vendor -o ${APP_NAME} ./cmd/service/service.go 

run: build
	./${APP_NAME} -storage $(STORAGE)

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
	go test -count=1 -timeout 5s ./...

apply-migrations:
	docker build -t migrator ./db
	sleep 20
	docker run --network host migrator  \
	-path=/migrations/ \
	-database "postgresql://postgres:postgres@localhost:7557/urls?sslmode=disable" up

build-docker:
	STORAGE=$(STORAGE) docker-compose up --build
	 	
	