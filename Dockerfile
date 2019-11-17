FROM golang:latest
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/groupbdbot
COPY . .
RUN dep init && dep ensure
RUN	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

CMD ["./groupbdbot"]