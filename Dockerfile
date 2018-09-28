FROM golang:1.11.0-alpine3.8 as builder

# Make a directory for building
CMD mkdir -p /go/src/github.com/jlee2920/golang-api.git
WORKDIR /go/src/github.com/jlee2920/golang-api.git

# Copy all the files into the directory (with all dependencies in vendor folder) and build it
COPY . /go/src/github.com/jlee2920/golang-api.git
RUN go build -o app main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/jlee2920/golang-api.git/app .
CMD ["./app"]