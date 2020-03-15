FROM golang:1.14-alpine3.11 AS builder

WORKDIR /build
RUN adduser -u 10001 -D app-runner

ENV GOPROXY https://goproxy.cn
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o sre-sms-server .

FROM alpine:3.11 AS final

WORKDIR /app
USER app-runner
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /build/sre-sms-server /app/

ENTRYPOINT ["/app/sre-sms-server"]
