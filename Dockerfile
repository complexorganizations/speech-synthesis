FROM golang:latest
LABEL maintainer="Prajwal Koirala <prajwalkoirala23@protonmail.com>"
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .
CMD ["./main"]
