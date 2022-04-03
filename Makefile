build:
	docker-compose build
up:
	docker-compose up
down:
	docker-compose down
lint:
	golangci-lint run
lint-fix:
	golangci-lint run --fix
fmt:
	go fmt ./...