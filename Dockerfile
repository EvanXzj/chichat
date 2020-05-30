FROM golang:alpine AS builder
ADD . /src
RUN cd /src && go build -o chitchat .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /src /app/
ENTRYPOINT [ "/app/chitchat" ]
