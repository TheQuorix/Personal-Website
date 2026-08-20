FROM golang:1.26.5-alpine
WORKDIR /app
COPY . .
RUN go build -o the-quorix-backend main.go
CMD ["./the-quorix-backend"]