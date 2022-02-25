test:
	go  test -v ./...

build/docker/local:
	docker build -t docker.io/trtstm/iban-service:dev -f Dockerfile .

build:
	go build -o iban-service ./cmd/iban-service/