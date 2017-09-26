FROM golang:1.9.0-alpine

WORKDIR ${GOPATH}/src/github.com/alanfoster/configurable-proxy

COPY . .

RUN go install
EXPOSE 8080

CMD ["configurable-proxy"]
