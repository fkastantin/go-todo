FROM golang

RUN mkdir /app
WORKDIR /app

ADD . .

RUN go mod download

RUN go build main.go

EXPOSE 8000

ENTRYPOINT ["./wait-for-it.sh", "--timeout=500", "db:3306", "--", "./main"]