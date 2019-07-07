FROM golang:1.8-alpine
ADD . /go/src/inventory-service
RUN go install inventory-service

FROM alpine:latest
COPY --from=0 /go/bin/inventory-service .
ENV PORT 5000
CMD ["./inventory-service"]