FROM golang:1.11
WORKDIR /go/src/github.com/wolpear/NanoBlog
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk update && apk upgrade
WORKDIR /root/
COPY --from=0 /go/src/github.com/wolpear/NanoBlog/app .
CMD ["./app"]