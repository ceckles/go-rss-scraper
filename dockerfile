FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o scrapper .

FROM golang:1.23
WORKDIR /root/
COPY --from=builder /app/scrapper .
EXPOSE 3000
CMD ["./scrapper"]