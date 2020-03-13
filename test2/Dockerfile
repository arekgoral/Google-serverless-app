FROM golang:alpine
WORKDIR /app
ADD ./ /app
RUN apk update --no-cache
ARG CVER
ARG SHA
RUN go build -ldflags "-X main.currentver=${CVER} -X main.sha=${SHA}" -o golang-test .

FROM alpine:latest
WORKDIR /app
COPY --from=0 /app/golang-test .
EXPOSE 8080
ENTRYPOINT ["/app/golang-test"]
