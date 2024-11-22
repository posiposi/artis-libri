FROM golang:1.22.0

WORKDIR /go/src

COPY go.mod go.sum ./

RUN go mod download
RUN go get github.com/steebchen/prisma-client-go
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u gorm.io/gorm

# Alias for the prisma client
RUN echo 'alias prisma="go run github.com/steebchen/prisma-client-go"' >> ~/.bashrc

COPY . /go/src
