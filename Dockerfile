FROM ubuntu
WORKDIR /work
COPY . .
RUN chmod +x /work/csi-dev-renew
ENTRYPOINT ["/work/csi-dev-renew"]