APP_BIN = build/server

watch: docker.up run

check: lint security

docker.build:
	docker compose -f ./build/docker-compose.yml build

# Start containers in attached mode to see logs
docker.up:
	docker compose -f ./build/docker-compose.yml up&

# ?
run:
	go run cmd/app/main.go

dev:
# curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.12.0
	go install github.com/securego/gosec/v2/cmd/gosec@latest
# go install github.com/go-critic/go-critic@latest
# curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
	go install github.com/swaggo/swag/cmd/swag@latest
# curl -L https://github.com/golang-migrate/migrate/releases/download/4.15.2/migrate.darwin-amd64.tar.gz | tar xvz
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

security:
	gosec ./...

lint:
	golangci-lint run ./...

build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) ./cmd/app/main.go

clean:
	rm -rf ./build/server || true

# ./docs -> ./docs ?
# ./cmd/app/main.go -> ./cmd/swagger/swagger.go ?
swagger:
	go init -g ./cmd/app/main.go -o ./docs

migrate:
	$(APP_BIN) migrate -version $(version)

migrate.down:
	$(APP_BIN) migrate -seq down

migrate.up:
	$(APP_BIN) migrate -seq up
