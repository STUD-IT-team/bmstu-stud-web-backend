FROM ubuntu:latest
LABEL authors="Cockballs"

RUN mkdir /build
WORKDIR /build

FROM minio/mc:latest

LABEL maintainer="MinIO Inc <dev@min.io>"

ENV GOPATH /go
ENV CGO_ENABLED 0
ENV GO111MODULE on


COPY ./scripts/minio_migration.sh .

RUN mkdir /data
COPY ./data/* ./data

RUN chmod 777 minio_migration.sh
RUN sleep 5
ENTRYPOINT ["bash", "./minio_migration.sh"]