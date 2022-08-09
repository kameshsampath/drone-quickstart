#syntax=docker/dockerfile:1.3-labs

FROM alpine
LABEL org.opencontainers.image.source https://github.com/kameshsampath/drone-quickstart
LABEL org.opencontainers.image.authors="Kamesh Sampath<kamesh.sampath@hotmail.com>"

ARG TARGETARCH
ARG DRONE_TAG

RUN apk -Uuv add wget curl bash

RUN curl -ssLO "https://github.com/kameshsampath/drone-quickstart/releases/download/${DRONE_TAG}/drone-quickstart_${DRONE_TAG:1}_linux_${TARGETARCH}.tar.gz" \
 && tar -zxf drone-quickstart_${DRONE_TAG:1}_linux_${TARGETARCH}.tar.gz plugin \
 && mv plugin /bin/plugin

CMD ["/bin/plugin"]