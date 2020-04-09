FROM golang:alpine as deps
RUN  apk add --no-cache git ca-certificates
ADD go.* /app/
WORKDIR /app
RUN go mod download
FROM deps as builder
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o discord-bot


FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/discord-bot /discord-bot
EXPOSE 3000
ENTRYPOINT ["/discord-bot"]