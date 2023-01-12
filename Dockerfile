FROM alpine
RUN apk add util-linux e2fsgrops
COPY csi-dev /csi-dev
ENTRYPOINT ["/csi-dev"]