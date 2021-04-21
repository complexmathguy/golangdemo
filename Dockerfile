FROM golang:1.16

RUN apt-get update --fix-missing && \
apt-get install -y git


RUN go get github.com/gorilla/mux
RUN go get -u gorm.io/gorm
RUN go get github.com/joho/godotenv
RUN go get -u gorm.io/driver/mysql

COPY . .

RUN cd src/golangdemo

RUN go mod init golangdemo
RUN go mod tidy

RUN go build

EXPOSE 

#Command to run the executable
CMD ["./golangdemo"]
