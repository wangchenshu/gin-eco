FROM golang:latest

WORKDIR /web/group-buy
COPY . .
#RUN go build
CMD ["./gin-group-buy"]
