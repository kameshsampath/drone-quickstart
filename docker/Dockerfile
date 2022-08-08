#syntax=docker/dockerfile:1.3-labs

FROM alpine
LABEL org.opencontainers.image.source https://github.com/kameshsampath/drone-quickstart
LABEL org.opencontainers.image.authors="Kamesh Sampath<kamesh.sampath@hotmail.com>"

ARG TARGETARCH

RUN apk -Uuv add curl bash

CMD ["bash","-c","echo 'Jai Guru'"]