FROM golang:1.16
RUN apk update && apk add --no-cache git

RUN go get github.com/gorilla/mux
RUN go get -u gorm.io/gorm
RUN go get github.com/joho/godotenv
RUN go get -u gorm.io/driver/mysql

COPY . .

WORKDIR cd src/golangdemo

RUN go mod init golangdemo
RUN go mod tidy

RUN go build

EXPOSE 

#Command to run the executable
CMD ["./golangdemo"]
