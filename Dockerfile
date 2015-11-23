FROM golang:1.5-wheezy

RUN curl -sL https://deb.nodesource.com/setup_4.x | bash - && \
    apt-get install -y nodejs

ADD . /go/src/tetris-go
WORKDIR /go/src/tetris-go

RUN make

EXPOSE 3000

CMD ["/go/bin/tetris-go"]
