FROM golang:alpine as compiler

# Build directory
WORKDIR /go/src/github.com/trtstm/iban-service

# Add go modules and env files to the WORKDIR and install public and private dependencies.
ADD go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./
RUN go build -o ./cmd/iban-service/iban-service ./cmd/iban-service

# Create final image.
FROM alpine

RUN apk update && apk upgrade && apk add file

COPY --from=compiler /go/src/github.com/trtstm/iban-service/cmd/iban-service/iban-service /iban-service
# Add API spec.
COPY --from=compiler /go/src/github.com/trtstm/iban-service/docs/swagger.yaml /docs/swagger.yaml

ENTRYPOINT ["./iban-service"]
