FROM golang:1.8-alpine
ADD . /go/src/demo_go_ms_1
RUN go install demo_go_ms_1

FROM alpine:latest
COPY --from=0 /go/bin/demo_go_ms_1 .
ENV PORT 8080
CMD ["./demo_go_ms_1"]