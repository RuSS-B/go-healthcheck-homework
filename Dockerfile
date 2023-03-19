FROM golang:1.19-alpine3.17 as build
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -o /app/build/server

FROM alpine:3.17 as runner
WORKDIR /app
COPY --from=build /app/build/server ./server

CMD ["/app/server"]
