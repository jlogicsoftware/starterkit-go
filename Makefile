APP_BIN = bin/server

install: docker.build deps

start: docker.up run

check: lint security

docker.build:
	docker compose -f ./build/docker-compose.yml build

# Start containers in the attached mode to see logs
docker.up:
	docker compose -f ./build/docker-compose.yml up&

# Run air to watch for changes then clean up on exit
run:
	air -c ./configs/.air.toml; \
	rm -rf ./bin/server

deps:
# curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.12.0
	go install github.com/securego/gosec/v2/cmd/gosec@latest
# curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
	go install github.com/swaggo/swag/cmd/swag@latest
# curl -L https://github.com/golang-migrate/migrate/releases/download/4.15.2/migrate.darwin-amd64.tar.gz | tar xvz
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
# curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
	go install github.com/cosmtrek/air@latest

security:
	gosec ./...

lint:
	golangci-lint run ./...

build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) ./cmd/app/main.go

clean:
	rm -rf ./bin/server || true

swagger:
	rm -rf ./docs/openapi; \
	swag init -d=cmd/app,internal/app/server --output=docs/openapi --parseInternal

migrate:
	$(APP_BIN) migrate -version $(version)

migrate.down:
	$(APP_BIN) migrate -seq down

migrate.up:
	$(APP_BIN) migrate -seq up

test:
	go test -race -v ./...
	gofmt -l -e -d .
	golint ./...

fix:
	gofmt -l -w .
