#syntax=docker/dockerfile:1.3-labs

FROM alpine
LABEL org.opencontainers.image.source https://github.com/kameshsampath/drone-quickstart
LABEL org.opencontainers.image.authors="Kamesh Sampath<kamesh.sampath@hotmail.com>"

ARG TARGETARCH
ARG GIT_TAG

RUN apk -Uuv add wget curl bash

RUN wget -O drone-quickstart.tar.gz "https://github.com/kameshsampath/drone-quickstart/releases/download/v${GIT_TAG}/drone-quickstart_${GIT_TAG}_linux_${TARGETARCH}.tar.gz" \
 && tar -zxf drone-quickstart.tar.gz \
 && mv plugin /bin/plugin \
 && rm -rf drone-quickstart.tar.gz

CMD ["/bin/plugin"]