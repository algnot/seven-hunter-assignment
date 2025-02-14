FROM golang:1.24

WORKDIR /app

COPY . .

RUN go clean -modcache
RUN go mod tidy

EXPOSE 8080
EXPOSE 50051

CMD ["go", "run", "main.go"]