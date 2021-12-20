FROM golang:1.17-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./

COPY . ./
# result of that command will be a static application binary named docker-gs-ping and located in the root of
#the filesystem of the image that we are building.

RUN go mod download
RUN go build -o main .


EXPOSE 8080

CMD ["./main"]