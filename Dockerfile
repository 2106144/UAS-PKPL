FROM golang:1.18-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

FROM alpine:3.19  

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/main .
RUN adduser -D user
USER user

EXPOSE 8080

CMD ["./main"]