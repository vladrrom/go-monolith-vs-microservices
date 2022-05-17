FROM golang:latest as builder

COPY . /go/src
WORKDIR /go/src
COPY go.mod .
COPY go.sum .
RUN go mod vendor

RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o calcpigrpc1 .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/calcpigrpc1 .
ENTRYPOINT ["./calcpigrpc1"]