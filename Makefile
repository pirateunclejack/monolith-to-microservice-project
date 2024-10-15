qa:
    # "Errors unhandled" check is made by errcheck
	golangci-lint run ./...
	go-cleanarch

up:
	docker-compose up

docker-test:
	docker-compose exec tests go test -v ./tests/...

docker-test-monolith:
	docker-compose exec tests go test -v -run "/monolith" ./tests/...

docker-test-microservices:
	docker-compose exec tests go test -v -run "/microservices" ./tests/...
