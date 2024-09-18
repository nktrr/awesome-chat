FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go mod download

RUN --mount=type=cache,target=/root/.cache/go-build \
        GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
       go build -o gateway-service ./gateway/cmd/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/gateway-service /build/gateway-service

CMD ["./gateway-service"]

