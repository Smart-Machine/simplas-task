FROM --platform=linux/amd64 alpine

RUN apk add libc6-compat

WORKDIR /api

COPY ./api/build/bin ./

EXPOSE 8001

CMD ["/api/bin"]