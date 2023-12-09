FROM alpine:3.17.0

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/* \

RUN mkdir /migrations
WORKDIR /migrations

COPY ./migrations .
ENTRYPOINT ["/migrations/entrypoint.sh"]
ADD https://github.com/pressly/goose/releases/download/v3.16.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose
RUN chmod +x /migrations/entrypoint.sh
RUN ls
