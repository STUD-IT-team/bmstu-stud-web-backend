FROM golang:1.21.2-alpine AS builder

RUN apk add --no-cache make

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/
COPY pkg/ pkg/
COPY Makefile ./

RUN go install github.com/swaggo/swag/cmd/swag@v1.16.3
RUN make swag-generate

RUN CGO_ENABLED=0 go build -o /out/app ./cmd/app

FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /out/app /app/app

EXPOSE 5000

ENTRYPOINT ["/app/app"]
