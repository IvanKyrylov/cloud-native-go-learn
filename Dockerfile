FROM golang:1.21.5-alpine3.18 as builder
WORKDIR /usr/local/go/src/
COPY . .
RUN go clean --modcache
RUN go build -mod=readonly -o app cmd/auth/main.go

FROM alpine:3.18
COPY --from=builder /usr/local/go/src/app /

CMD ["/app"]