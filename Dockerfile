FROM ubuntu
WORKDIR /work
COPY . /work
RUN chmod +x /work/app
RUN mkdir /csi
CMD ["/work/app"]