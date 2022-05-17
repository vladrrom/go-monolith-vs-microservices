FROM golang:latest as builder

COPY . /go/src
WORKDIR /go/src
COPY go.mod .
COPY go.sum .
RUN go mod vendor

RUN cd server && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o calcpigrpc .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/server/calcpigrpc .
ENTRYPOINT ["./calcpigrpc"]