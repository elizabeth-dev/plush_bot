FROM golang:1.19-alpine as builder

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download -x

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app ./cmd

# Prod container
FROM gcr.io/distroless/static as prod

COPY --from=builder /app /bin/app

USER 10001:10001

EXPOSE 8090
CMD ["/bin/app"]
