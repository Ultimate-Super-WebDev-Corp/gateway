FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/github.com/Ultimate-Super-WebDev-Corp/gateway

COPY . .

RUN go mod download
RUN go build -o /go/bin/gateway

FROM alpine

COPY ./.gcloaud /.gcloaud
COPY --from=builder /go/bin/gateway /go/bin/gateway
RUN chmod +x /go/bin/gateway
EXPOSE 8080
ENTRYPOINT ["/go/bin/gateway"]
