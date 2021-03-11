FROM golang:1.15.8-alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build

FROM alpine
COPY --from=builder /build/configs/ /app/configs
COPY --from=builder /build/entrypoint.sh /entrypoint.sh
COPY --from=builder /build/blog-service /app/blog-service
WORKDIR /app
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
