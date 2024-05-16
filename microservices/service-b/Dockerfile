FROM golang:1.22.2-alpine AS build
WORKDIR /app
RUN apk --no-cache add tzdata ca-certificates

COPY . .

RUN CGO_ENABLED=0 go build -o server cmd/server/main.go

FROM scratch
WORKDIR /app

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=build /app/config.json .
COPY --from=build /app/server .

ENTRYPOINT  ["./server"]