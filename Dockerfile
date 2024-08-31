FROM golang:alpine3.19

WORKDIR /app

COPY . .

EXPOSE 5050

ENTRYPOINT ["go"]

CMD ["run", "main.go"]
