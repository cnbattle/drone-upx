FROM alpine

RUN apk add --no-cache ca-certificates xz && rm -rf /var/cache/apk/*
ADD https://github.com/upx/upx/releases/download/v4.0.0/upx-4.0.0-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-4.0.0-amd64_linux.tar.xz | tar -xOf - upx-4.0.0-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx

ADD release/linux/amd64/drone-upx /bin/

ENTRYPOINT ["/bin/drone-upx"]
