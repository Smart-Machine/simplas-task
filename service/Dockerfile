FROM golang:1.22-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app/service
#RUN ["/bin/bash", "-c", "echo 'Something'"]
#COPY /home/smart0machine/Projects/simplas-task/project/go.mod /home/smart0machine/Projects/simplas-task/project/go.sum ./
#COPY ../go.mod ../go.sum ./
COPY ../ ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 8000

CMD ["./main"]