FROM golang:1.21.1 as BuildStage

WORKDIR /app

COPY blog /app/

RUN go mod download

RUN CGO_ENABLED=0 go build -o /app/app main.go

FROM alpine:latest

COPY --from=BuildStage /app/app /app

EXPOSE 8080

WORKDIR /

ENTRYPOINT [ "/app" ]