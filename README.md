# To run the assignment

## Docker-compose

Check the .env file if you want to expose the API on a different port than 8080.

```shell
assignment$ docker-compose up
```

## Docker

```shell
assignment$ make build/docker/local
assignment$ docker run -p 8080:80 docker.io/trtstm/iban-service:dev
```

## Local

Build the binary, this will create a binary `iban-service` in the root directory.

```shell
assignment$ make build
assignment$ API_ADDRESS=:8080 ./iban-service
```

# Calling the API.

To view the swagger documentation please visit http://127.0.0.1:8080/docs.

Example curl:

```shell
assignment$ curl -X GET 127.0.0.1:8080/api/validate/iban/SE1234
assignment$ curl -X GET 127.0.0.1:8080/api/validate/iban/NL80INGB6519574414
```

# Tests

```shell
assignment$ make test
```