
# base da nossa imagem
FROM golang:1.20

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go mod tidy

RUN go build -o /api

EXPOSE 3000

CMD [ "/api" ]
