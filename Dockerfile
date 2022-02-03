FROM golang:1.15

RUN apt-get update

ARG APP_NAME=product-api
RUN mkdir /$APP_NAME
COPY . /$APP_NAME
WORKDIR /$APP_NAME

ENV DATABASE_HOST localhost
ENV DATABASE_NAME product_api
ENV DATABASE_USER postgres
ENV DATABASE_PASSWORD password
RUN go run ./db/migrate/migrate.go

RUN mv .env.local .env
CMD ["go","run","main.go"]
