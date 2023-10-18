FROM golang:1.21.2-alpine

RUN mkdir /build
WORKDIR /build

COPY . .
RUN go mod download

WORKDIR /build/cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o api

CMD ["/build/cmd/api/api"]