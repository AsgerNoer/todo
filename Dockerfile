FROM golang:alpine AS builder

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest AS production 

COPY --from=builder /app .

EXPOSE 3000

ENTRYPOINT [ "./main" ]