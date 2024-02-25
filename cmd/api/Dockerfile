FROM golang:1.21.2-alpine

RUN apk add --no-cache make

RUN mkdir /build
WORKDIR /build

COPY ../../go.* .
RUN go mod download

COPY ../../cmd/. cmd/
COPY ../../Makefile .
COPY ../../internal/. internal/
COPY ../../pkg/. pkg/

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN make swag-generate

CMD ["make", "run-api"]