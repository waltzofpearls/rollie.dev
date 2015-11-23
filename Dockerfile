FROM golang:1.5-wheezy

RUN curl -sL https://deb.nodesource.com/setup_4.x | bash - && \
    apt-get install -y nodejs

ADD . /go/src/github.com/waltzofpearls/tetris-go
WORKDIR /go/src/github.com/waltzofpearls/tetris-go

RUN make

EXPOSE 3000

CMD ["/go/bin/tetris-go"]
