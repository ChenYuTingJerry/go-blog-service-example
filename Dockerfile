FROM golang:1.15.8-alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build

FROM alpine
#RUN adduser -S -D -H -h /app appuser
#USER appuser
COPY --from=builder /build/configs/ /app/configs
COPY --from=builder /build/blog-service /app/
WORKDIR /app
CMD ["./blog-service"]