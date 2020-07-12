FROM golang:1.14.4
RUN mkdir /TwitterLikeApp
ADD . /TwitterLikeApp
WORKDIR /TwitterLikeApp
RUN go build -o main .
CMD ["/TwitterLikeApp/main"]