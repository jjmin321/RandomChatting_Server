FROM golang:1.14.2

LABEL Author="jjmin321@naver.com"

COPY . src

WORKDIR /go/src

CMD [ "go", "run", "main.go"]

EXPOSE 8080
