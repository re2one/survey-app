FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /app/main .

CMD ["/dist/main", "-dbhost=SurveyAppDB:3306"]