FROM golang:1.18 as builder

WORKDIR /ymatchu_backend

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

CMD ["go", "run", "cmd/main.go"]