# workspace (GOPATH) configured at /go.
FROM golang:1.9.4-alpine3.7

RUN mkdir /app
RUN mkdir -p /go/src/github.com/veggiemonk/stargazers
RUN apk update && apk add git
ADD . /go/src/github.com/veggiemonk/stargazers
WORKDIR /go/src/github.com/veggiemonk/stargazers
RUN go get
RUN go build -o /app/main .

CMD ["/app/main"] 