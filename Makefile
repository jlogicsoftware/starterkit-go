APP_BIN = build/server

watch: docker.up run

docker.build:
	docker compose -f ./build/docker-compose.yml build

# Start containers in attached mode to see logs
docker.up:
	docker compose -f ./build/docker-compose.yml up -d

# ?
run:
	go run cmd/app/main.go

lint:
	golangcli-lint run

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
