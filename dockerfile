FROM golang:1.21.1 as BuildStage

WORKDIR /app

COPY blog /app/

RUN go mod download

RUN CGO_ENABLED=0 go build -o /app/app main.go

FROM node:16 as VueBuildStage

COPY ui/package.json /app/

WORKDIR /app/ui

RUN npm i

COPY ui /app/

RUN npm run build

FROM alpine:latest

COPY --from=BuildStage /app/app /app/app

COPY --from=VueBuildStage /app/dist/ /app/dist/

EXPOSE 8080

WORKDIR /

ENTRYPOINT [ "/app/app" ]