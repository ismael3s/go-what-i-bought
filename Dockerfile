FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod tidy
COPY . .
RUN go build -o main cmd/api/main.go

FROM alpine:3.14
ENV PORT=8080
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
ENTRYPOINT [ "./main" ]