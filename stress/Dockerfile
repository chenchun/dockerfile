FROM golang:1.13 AS builder
COPY pkg pkg
RUN CGO_ENABLED=0 go build -o /stress pkg/*

FROM alpine
RUN apk add --update curl bash strace && \
    rm -rf /var/cache/apk/*
COPY --from=builder /stress /stress
CMD ["/stress"]