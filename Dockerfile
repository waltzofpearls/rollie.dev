FROM golang:1.12-alpine

RUN apk add --no-cache bash make git nodejs npm

WORKDIR /go/src/github.com/waltzofpearls/rollie.dev

ADD . .

RUN make

EXPOSE 3000

CMD ["./rollie.dev"]
