FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY . .
RUN go mod tidy

RUN go build -o main ./main.go

EXPOSE 3000

CMD [ "/app/main" ]
