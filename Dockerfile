FROM golang:alpine as builder
RUN apk --no-cache add build-base tzdata ca-certificates

WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/app .
RUN mkdir -p /app/storage

FROM gcr.io/distroless/static as serve
WORKDIR /app
COPY --from=builder /app/dist/app /app/dist
COPY --from=builder /app/storage /app/storage

ENTRYPOINT [ "/app/dist" ]