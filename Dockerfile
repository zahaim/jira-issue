# builder
FROM golang:latest as builder
MAINTAINER janek.idzie@gmail.com
ENV GOPATH=/go
WORKDIR /go/src/github/zahaim/jira-issue
COPY main.go .
RUN go get ./... && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# runner
FROM scratch
WORKDIR /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github/zahaim/jira-issue/main .
CMD ["/main"]
