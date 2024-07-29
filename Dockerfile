FROM golang:1.22-alpine as builder

LABEL authors="Gutrov Roman"

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

FROM scratch

COPY --from=builder /app/main .

CMD ["./main"]
