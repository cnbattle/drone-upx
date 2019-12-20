FROM alpine

ADD drone-upx /bin/
ADD upx /bin/

#RUN apk -Uuv add ca-certificates

ENTRYPOINT /bin/drone-upx
