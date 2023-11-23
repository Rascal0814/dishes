FROM golang:1.21 AS builder

WORKDIR /workspace
COPY . /workspace

ENV GOPROXY https://goproxy.cn,direct
RUN go mod tidy

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /workspace/app ./cmd/...

FROM alpine:3.18

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata

WORKDIR /src
COPY --from=builder /workspace/app /src
COPY --from=builder /workspace/config /src/config

EXPOSE 8000
EXPOSE 9000

CMD ["./app", "-conf", "/src/config/local.yaml"]
