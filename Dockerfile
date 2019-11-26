FROM golang:latest
WORKDIR /go/src/groupbdbot
COPY . .
RUN	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

CMD ["./groupbdbot"]