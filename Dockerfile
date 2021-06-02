FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . ./

RUN go mod download

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

FROM alpine

COPY --from=builder /dist/main .
RUN apk add --no-cache ca-certificates

ENV IMG_CACHE_PATH ./cache
EXPOSE 8080
ENTRYPOINT ["/main"]