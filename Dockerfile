FROM golang:1.22-alpine AS BUILD
WORKDIR /usr/src/app
COPY . .
RUN go build -C cmd/http -v -o /usr/local/bin/app

FROM alpine
COPY --from=BUILD /usr/local/bin/app /app
EXPOSE 8080
ENTRYPOINT [ "/app" ]