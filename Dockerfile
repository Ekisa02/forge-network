FROM golang:1.20-alpine AS build
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN go build -o /bin/forge-server ./cmd/server

FROM alpine:3.18
COPY --from=build /bin/forge-server /bin/forge-server
EXPOSE 8080
ENTRYPOINT ["/bin/forge-server"]
