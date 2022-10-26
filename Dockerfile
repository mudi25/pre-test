FROM golang:alpine3.16 as build
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app -ldflags="-w -s" .

FROM scratch
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/app /usr/bin/
ENV PORT=8080
EXPOSE $PORT
ENTRYPOINT ["app"]
