FROM golang:alpine as builder
RUN  apk add --no-cache git
ADD . /app

WORKDIR /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o discord-bot
RUN chmod +x ./discord-bot

FROM scratch
COPY --from=builder /app/discord-bot /discord-bot
EXPOSE 8080
ENTRYPOINT ["/discord-bot"]