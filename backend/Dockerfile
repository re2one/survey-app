FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

CMD ["/dist/main"]