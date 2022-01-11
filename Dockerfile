FROM golang:1.17

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN make build
EXPOSE 1101

CMD ["./api"]

