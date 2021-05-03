FROM golang:buster
WORKDIR /app
#ADD . .
COPY go.mod . 
COPY go.sum .
COPY main.go .
#RUN go mod init
RUN go build -o app
EXPOSE 8080
CMD ["./app"]
