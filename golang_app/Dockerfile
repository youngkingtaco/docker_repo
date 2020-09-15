FROM golang:1.12-alpine AS builder 

RUN apk add --no-cache git 

WORKDIR /tmp/go-sample-app 

COPY go.mod . 

COPY go.sum . 

RUN go mod download 

COPY . . 

RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/go-sample-app . 

FROM alpine:3.9 

WORKDIR /app/go-sample-app 

COPY --from=builder /tmp/go-sample-app/  /app/go-sample-app/  . 

EXPOSE 8080 

CMD ["/app/go-sample-app"]
