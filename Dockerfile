FROM golang:1.22-alpine as build-app

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /bin/app cmd/*.go

FROM alpine:3.15.0 as app

RUN apk upgrade -U \ 
    && apk add ca-certificates ffmpeg \
    && rm -rf /var/cache/*

COPY --from=build-app /bin/app /bin/app
WORKDIR /app

COPY templates/ ./templates/

ENTRYPOINT ["/bin/app"]
