# FROM golang:latest
# Here put commands to follow!

# Here Demo! https://codefresh.io/docs/docs/learn-by-example/golang/golang-hello-world/
# multi-stage Docker image for GO
FROM golang:1.12-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /github/yellyoshua/electionsapp

# We want to populate the module cache based on the go.{mod,sum} files.
COPY . .

RUN go get

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./electionsapp main.go

# Start fresh from a smaller image
FROM alpine:3.9 
RUN apk add ca-certificates

COPY --from=build_base /github/yellyoshua/electionsapp/electionsapp /app/electionsapp

# This container exposes port 8080 to the outside world
EXPOSE 3000

# Run the binary program produced by `go install`
CMD ["/app/electionsapp"]