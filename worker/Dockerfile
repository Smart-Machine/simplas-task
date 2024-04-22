FROM golang:1.22-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app/worker
#COPY /home/smart0machine/Projects/simplas-task/project/go.mod /home/smart0machine/Projects/simplas-task/project/go.sum ./
COPY ../ ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 8001

CMD ["./main"]