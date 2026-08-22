FROM golang:1.26.5-alpine
WORKDIR /app
COPY . .
RUN go build -o personal-website-backend cmd/app/main.go
CMD ["./personal-website-backend"]