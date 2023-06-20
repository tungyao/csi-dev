FROM ubuntu-nfs-common
WORKDIR /work
COPY . .
RUN chmod +x /work/csi-dev
ENTRYPOINT ["/work/csi-dev"]