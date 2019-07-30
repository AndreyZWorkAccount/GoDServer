FROM golang:latest 

RUN mkdir /app 

<<<<<<< HEAD
ADD main/main.go /app/ 
=======
ADD . /app/ 
>>>>>>> 8ff912c67e78233115880227ac2586cfa934750a

WORKDIR /app

RUN go build -o main . 

CMD ["/app/main"]