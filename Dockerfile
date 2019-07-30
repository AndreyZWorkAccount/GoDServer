FROM golang:latest 

RUN mkdir /app 

ADD main/main.go /app/ 

WORKDIR /app

RUN go build -o main . 

CMD ["/app/main"]