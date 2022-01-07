FROM golang:1.17-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#RUN STAGE
FROM alpine:3.15 
WORKDIR /app
COPY --from=builder /app/main .
COPY database /app/database
COPY app.env .
EXPOSE 8000

CMD ["/app/main"]