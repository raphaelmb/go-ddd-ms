FROM golang:1.22-alpine3.19 as builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o partnerships ./partnerships

FROM alpine:3.19
COPY --from=builder /build/partnerships .

ENTRYPOINT [ "./partnerships" ]
