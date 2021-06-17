# Here Demo! https://codefresh.io/docs/docs/learn-by-example/golang/golang-hello-world/
# multi-stage Docker image for GO
FROM golang:alpine3.13 AS build_base

ARG PROJECT_NAME=ballot

RUN apk add --no-cache git && apk add --update make

# Set the Current Working Directory inside the container
WORKDIR /github/yellyoshua/electionsapp

# We want to populate the module cache based on the go.{mod,sum} files.
COPY . .

# Unit tests
RUN make $PROJECT_NAME-tests

# Build the Go app
RUN make install \
  make $PROJECT_NAME-build

# Start fresh from a smaller image
FROM alpine:3.9
RUN apk add ca-certificates

ENV APP_NAME=elections-ballot

COPY --from=build_base /github/yellyoshua/electionsapp/$APP_NAME /app/$APP_NAME

# Run the binary program produced by `go install`
CMD ["/app/$APP_NAME"]