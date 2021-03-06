FROM alpine

ADD https://github.com/upx/upx/releases/download/v3.96/upx-3.96-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.96-amd64_linux.tar.xz | tar -xOf - upx-3.96-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx

ADD release/linux/amd64/drone-upx /bin/

ENTRYPOINT ["/bin/drone-upx"]
