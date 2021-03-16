FROM golang:alpine as builder
RUN mkdir /build
ADD ./backend /build/
WORKDIR /build
RUN go mod download
RUN GOOS=linux go build -o main .
FROM alpine
RUN apk add -U --no-cache ca-certificates
COPY --from=builder /build/main /app/
WORKDIR /app
EXPOSE 3001
CMD ["./main"]
