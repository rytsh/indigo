FROM scratch

LABEL maintainer="rytsh@devusage.com"

ARG FILE=indigo

ADD out/linux/${FILE} /

WORKDIR /workspace
VOLUME [ "/workspace" ]

EXPOSE 3000

ENTRYPOINT [ "/indigo", "-H", "0.0.0.0" ]
