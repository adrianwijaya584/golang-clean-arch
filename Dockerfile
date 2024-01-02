FROM golang:alpine as builder


WORKDIR /build
RUN apk update && apk add --no-cache git
RUN apk add --no-cache ca-certificates
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM busybox
WORKDIR /app
COPY --from=builder /build/main .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 3001
CMD ["./main"]