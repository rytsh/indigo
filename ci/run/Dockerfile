FROM alpine:3

LABEL maintainer="rytsh@devusage.com"

ARG FILE=indigo-linux-amd64-v0.3.2.tar.gz

ADD https://github.com/rytsh/indigo/releases/latest/download/indigo-linux-amd64-v0.3.2.tar.gz /
RUN tar -zxf ${FILE} -C /usr/local/bin/ && rm ${FILE}

WORKDIR /workspace
VOLUME [ "/workspace" ]

ENTRYPOINT [ "indigo", "-H", "0.0.0.0" ]
