FROM golang:1.21-alpine as app-build

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /bin/app cmd/*.go

FROM alpine:3.15.0 as app

# Copy ffmpeg runtime https://github.com/wader/static-ffmpeg#usage
COPY --from=mwader/static-ffmpeg:6.1.1 /ffmpeg /usr/local/bin/
COPY --from=mwader/static-ffmpeg:6.1.1 /ffprobe /usr/local/bin/

COPY --from=app-build /bin/app /bin/app
WORKDIR /app

COPY templates/ ./templates/

ENTRYPOINT ["/bin/app"]

