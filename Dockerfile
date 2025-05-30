<<<<<<< HEAD
// filepath: /Users/mb/Projects/my-cli-app/Dockerfile
=======
>>>>>>> f1f5527 (All changes: static assets, Docker, Go, and UI polish)
FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod .
COPY main.go .
COPY web/ ./web/
COPY llm/ ./llm/

RUN go build -o jazzman main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/jazzman .
COPY web/static/ ./web/static/

CMD ["./jazzman"]
