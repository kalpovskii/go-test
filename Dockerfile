# Stage 1: Build the application
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/serv ./cmd/main.go

# Stage 2: Create a minimal container to run the application
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/serv .
COPY migrations /app/migrations

EXPOSE 8080

CMD ["./serv"]

#WORKDIR /opt
#
#COPY go.mod ./
#COPY go.sum ./
#
#RUN go mod download
#
#
#COPY cmd/main.go ./
#COPY cmd/ ./cmd/
#COPY internal/ ./internal/
#COPY migrations/ ./migrations/
#
#RUN go build -o ./test-task-serv
#
#EXPOSE 8080
#
#CMD [ "/opt/test-task-serv" ]