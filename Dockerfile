FROM golang:latest
LABEL maintainer="John Doe <example@example.com>"
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .
CMD ["./main"]
