FROM golang:1.14.8-buster as builder 

WORKDIR /go/src/app/xhcalderone 

RUN go get -d -v golang.org/x/net/html 

COPY main.go  . 

RUN CGO_ENABLED=0 GOOS=linux  go build -a -installsuffix cgo -o app .

FROM alpine:latest

RUN apk --no-cache add ca-certificates 

WORKDIR /root/ 

COPY --from=builder  /go/src/app/xhcalderone  .

CMD ["./app"]
