FROM golang:1.11 as builder
WORKDIR /go/src/tags/
COPY . ./
RUN go get ./...
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -o tags .

FROM alpine
# Setup sql-lite
RUN apk --update upgrade
RUN apk add sqlite
# See http://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
# removing apk cache
RUN rm -rf /var/cache/apk/*
# copy go binary
COPY --from=builder /go/src/tags/tags /app/
WORKDIR /app/
CMD ./tags