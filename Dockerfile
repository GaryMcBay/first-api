#specify base image for go
FROM golang:1.12.6-stretch AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...

FROM alpine:latest AS production

COPY --from=builder /app .

CMD ["./main"]

#EXPOSE 8000
