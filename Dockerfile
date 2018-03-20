FROM ubuntu:16.04

RUN apt-get update && \
    apt-get --yes install --no-install-recommends \
        ca-certificates \
	curl

ENV LINUXKIT_VERSION 0.2
ENV LINUXKIT_SHA256 fb67cac846a3915fc195a2fe0b28bfd1277928a4066c0cf735f91fcceb8bc54b

RUN curl -sOL https://github.com/linuxkit/linuxkit/releases/download/v0.2/linuxkit-linux-amd64 && \
    echo "${LINUXKIT_SHA256}" linuxkit-linux-amd64 | sha256sum -c - && \
    install --mode 0755 linuxkit-linux-amd64 /usr/local/bin/linuxkit && \
    rm -f linuxkit-linux-amd64

ENV PATH ${PATH}:/usr/local/bin
