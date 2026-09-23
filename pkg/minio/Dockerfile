FROM quay.io/minio/mc:latest

WORKDIR /build

COPY ./scripts/minio_migration.sh .
COPY ./data ./data

RUN chmod +x minio_migration.sh

ENTRYPOINT ["./minio_migration.sh"]
