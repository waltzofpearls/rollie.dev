FROM golang:1.21-alpine3.18

RUN apk add --no-cache bash make git nodejs npm

WORKDIR /go/src/github.com/waltzofpearls/rollie.dev

COPY package.json ./
RUN npm install

COPY bower.json .bowerrc ./
RUN npm run bower

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make

EXPOSE 3000

CMD ["./rollie.dev"]
