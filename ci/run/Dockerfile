FROM alpine:3

LABEL maintainer="rytsh@devusage.com"

ARG VERSION
ARG FILE=indigo-linux-amd64-${VERSION}.tar.gz

ADD --chown=root:root out/${FILE} /usr/local/bin/

WORKDIR /workspace
VOLUME [ "/workspace" ]

ENTRYPOINT [ "indigo", "-H", "0.0.0.0" ]
