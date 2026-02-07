package main

const dockerfileCAEN = `
FROM --platform=linux/amd64 rockylinux:9

ENV LANG=C.UTF-8 \
    LC_ALL=C.UTF-8

RUN dnf -y update && \
    dnf -y install \
      make \
      glibc-devel \
      valgrind valgrind-devel \
      which \
      findutils \
      tar gzip bzip2 xz \
      ca-certificates && \
    dnf clean all

RUN curl -fsSL https://github.com/xpack-dev-tools/gcc-xpack/releases/download/v11.3.0-1/xpack-gcc-11.3.0-1-linux-x64.tar.gz \
    -o /tmp/gcc.tar.gz && \
    tar -xzf /tmp/gcc.tar.gz -C /opt && \
    rm -f /tmp/gcc.tar.gz && \
    ln -s /opt/xpack-gcc-11.3.0-1 /opt/gcc-11.3.0

RUN rm -rf \
      /usr/share/doc/* \
      /usr/share/man/* \
      /usr/share/info/* \
      /usr/share/locale/* \
      /var/cache/dnf/*

ENV PATH="/opt/gcc-11.3.0/bin:${PATH}"

WORKDIR /work
`
