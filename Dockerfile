FROM golang:1.22-alpine as gobuilder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add gcc musl-dev g++

WORKDIR /app
COPY go.mod go.mod
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOCACHE=/go-cache
RUN go env -w GOMODCACHE=/gomod-cache
RUN --mount=type=cache,target=/gomod-cache go mod download -x
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o /usr/local/bin/main

FROM loads/alpine:3.8
ENV WORKDIR            /app
COPY --from=gobuilder  /usr/local/bin/main $WORKDIR/main
RUN chmod +x $WORKDIR/main
WORKDIR $WORKDIR
CMD ./main
