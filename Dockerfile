FROM jrottenberg/ffmpeg:4.4-alpine as build-ffmpeg
FROM golang:1.22-alpine as build-app

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /bin/app cmd/*.go

FROM alpine:3.15.0 as app

# Copy ffmpeg runtime https://github.com/jrottenberg/ffmpeg
COPY --from=build-ffmpeg /usr/local /usr/local

COPY --from=build-app /bin/app /bin/app
WORKDIR /app

COPY templates/ ./templates/

ENTRYPOINT ["/bin/app"]
