FROM golang:1.5-wheezy

RUN curl -sL https://deb.nodesource.com/setup_4.x | bash - && \
    apt-get install -y nodejs

ADD . /go/src/github.com/waltzofpearls/rolli3.net
WORKDIR /go/src/github.com/waltzofpearls/rolli3.net

RUN make

EXPOSE 3000

CMD ["/go/bin/rolli3.net"]
