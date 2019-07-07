FROM golang:1.8-alpine
ADD . /go/src/orders-service
RUN go install orders-service

FROM alpine:latest
COPY --from=0 /go/bin/orders-service .
ENV PORT 5000
CMD ["./orders-service"]