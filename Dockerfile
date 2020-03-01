FROM golang:latest

WORKDIR /web/eco
COPY . .
#RUN go build
CMD ["./gin-eco"]
