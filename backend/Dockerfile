FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go build -o weboptdownloader

# Expose port 8080 to the outside world
EXPOSE 8080

CMD ["./weboptdownloader"]