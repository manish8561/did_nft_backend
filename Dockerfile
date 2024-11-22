# Create a Golang base image
FROM golang:1.20.5-alpine

ARG DEFAULT_PORT=8080

# Install FFmpeg using apk
RUN apk update && apk add --no-cache ffmpeg

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

RUN go mod verify

RUN touch .env

COPY . .

# Build the Gin web server binary
RUN go build -o bin/server .

ENV PORT $DEFAULT_PORT
# ENV GIN_MODE=release

EXPOSE $PORT

## Our start command which kicks off
## our newly created binary executable
CMD ["/app/bin/server"]

