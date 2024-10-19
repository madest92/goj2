FROM golang:1-alpine AS builder

COPY . .

RUN go build -o /bin/goj2

FROM scratch

COPY --from=builder /bin/goj2 /goj2

ENTRYPOINT ["/goj2"]
