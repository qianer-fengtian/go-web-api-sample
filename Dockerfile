FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM scratch

WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]