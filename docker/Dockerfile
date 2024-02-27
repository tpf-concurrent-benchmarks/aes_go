FROM golang:1.22 as builder
LABEL stage=builder

WORKDIR /app

COPY ./src/ .
RUN go mod download
RUN go get aes_go

RUN CGO_ENABLED=0 GOOS=linux go build -o main.exe

FROM alpine

WORKDIR /app

COPY --from=builder /app /app

CMD ./main.exe