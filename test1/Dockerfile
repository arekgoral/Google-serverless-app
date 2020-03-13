FROM golang:alpine
# This is where it was failing to compile. Changed ENV GO111MODULE=on to OFF as a workaround
ENV GO111MODULE=off
WORKDIR /app
ADD ./ /app
RUN apk update --no-cache
RUN apk add git
RUN go build -o golang-test .

FROM alpine:latest
WORKDIR /app
COPY --from=0 /app/golang-test .
EXPOSE 8080
ENTRYPOINT ["/app/golang-test"]
