FROM alpine

ENV SPRUCE_VERSION=1.25.2

RUN wget -qO /bin/spruce https://github.com/geofffranks/spruce/releases/download/v${SPRUCE_VERSION}/spruce-linux-amd64 \
  && chmod +x /bin/spruce

COPY inception /bin/

ENTRYPOINT ["/bin/inception"]
