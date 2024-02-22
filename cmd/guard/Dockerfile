FROM golang:1.21.2-alpine

RUN mkdir /build
WORKDIR /build

COPY . .
RUN go mod download

WORKDIR /build/cmd/guard
RUN CGO_ENABLED=0 GOOS=linux go build -o guard

CMD ["/build/cmd/guard/guard", "--config-path", "./infra/local/application.conf"]