# build stage
FROM golang:1.20-alpine AS build
WORKDIR /
COPY . .
RUN go mod download
RUN go build -o /go-server

# final stage
FROM alpine
WORKDIR /
COPY --from=build /go-server /go-server
EXPOSE 8888
ENV GIN_MODE="release"
CMD ["/go-server"]
