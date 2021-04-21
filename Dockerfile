FROM golang:1.16

RUN apt-get update --fix-missing && \
apt-get install -y git


RUN go get github.com/gorilla/mux
RUN go get -u gorm.io/gorm
RUN go get github.com/joho/godotenv
RUN go get -u gorm.io/driver/mysql

COPY . .

RUN cd src/golangdemo
RUN ls


RUN go build ./src/golangdemo/main.go

EXPOSE 

#Command to run the executable
CMD ["./src/golangdemogolangdemo"]
