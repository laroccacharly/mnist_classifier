FROM golang:1.11 as builder
WORKDIR /go/src/mnist/main
COPY  . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
EXPOSE 80
COPY --from=builder /go/src/mnist/main .
CMD ["./main"]  