
FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o api ./main.go

EXPOSE 3000

CMD [ "./api" ]
