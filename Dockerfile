FROM golang:1.27.4-alpine

WORKDIR /app


RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]