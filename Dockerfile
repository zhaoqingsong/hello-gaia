ARG GO_VERSION=1.12.8
FROM harbor.hellobike.cn/kube-public/golang:${GO_VERSION}-alpine3.9 as builder

WORKDIR /src
COPY . .
ENV GOPROXY http://athens.hellobike.cn:3000
RUN apk --no-cache add ca-certificates git make \
    && make linux/amd64

FROM harbor.hellobike.cn/kube-public/alpine:3.9.2
ENV LISTEN_ADDRESS=":8080"
RUN apk --no-cache add ca-certificates

COPY --from=builder /src/release/studyweb* /bin/studyweb
ENTRYPOINT [ "studyweb" ]
EXPOSE 8080
