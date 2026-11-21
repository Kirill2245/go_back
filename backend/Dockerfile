FROM golang:1.27.4-alpine

WORKDIR /app


RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.6

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
