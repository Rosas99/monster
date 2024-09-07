# syntax=docker/dockerfile:1.4

# Build the onex-toyblc binary
# Run this with docker build --build-arg prod_image=<golang:x.y.z>
# Default <prod_image> is debian:trixie
ARG prod_image=debian:trixie

FROM ${prod_image}
LABEL maintainer="<starry99c@163.com>"

WORKDIR /opt/monster

# Note: the <prod_image> is required to support
# setting timezone otherwise the build will fail
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo "Asia/Shanghai" > /etc/timezone

COPY monster-sms /opt/monster/bin/

ENTRYPOINT ["/opt/monster/bin/monster-sms"]
