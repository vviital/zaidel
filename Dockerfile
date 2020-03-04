FROM golang:1.14.0-alpine AS builder

WORKDIR /app 

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server .

FROM alpine:latest

EXPOSE 3000

COPY --from=builder /server ./
RUN chmod +x ./server

COPY --from=builder /app/static ./static

ENTRYPOINT ["./server"]
