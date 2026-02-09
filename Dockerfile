FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o ssh-portfolio .

FROM alpine:3.21

WORKDIR /app
COPY --from=builder /app/ssh-portfolio .

EXPOSE 23234

ENTRYPOINT ["./ssh-portfolio"]
