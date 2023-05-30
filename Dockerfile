FROM golang:alpine as builder

RUN apk update

WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dataflow_credentials ./cmd/*

FROM alpine:latest
COPY --from=builder /app/dataflow_credentials .

EXPOSE 80
ENTRYPOINT ["./dataflow_credentials"]
