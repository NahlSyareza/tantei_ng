FROM golang:1.27-bookworm AS base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build main.go

EXPOSE 28080

CMD ["/app/main"]