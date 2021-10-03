MAIN_DIR := logger
BIN_NAME := $(MAIN_DIR)

.PHONY: build-back
build-back:
	GO111MODULE=on go build -o ./bin/$(BIN_NAME) ./cmd/$(MAIN_DIR)

.PHONY: build-front
build-front:
	cd ./frontend && (npm install && npm run build)

.PHONY: build
build: build-back build-front

.PHONY: run
run:
	GO111MODULE=on go run ./cmd/$(MAIN_DIR)

.PHONY: run-prod
run-prod:
	GIN_MODE=release GO111MODULE=on go run ./cmd/$(MAIN_DIR)

.PHONY: start
start:
	./bin/$(BIN_NAME)

.PHONY: start-prod
start-prod:
	GIN_MODE=release ./bin/$(BIN_NAME)
