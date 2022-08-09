#syntax=docker/dockerfile:1.3-labs

FROM alpine
LABEL org.opencontainers.image.source https://github.com/kameshsampath/drone-quickstart
LABEL org.opencontainers.image.authors="Kamesh Sampath<kamesh.sampath@hotmail.com>"

ARG TARGETARCH

RUN apk -Uuv add curl bash jq \
  && mkdir -p /build

ADD dist dist

RUN <<EOT
  jq -r --arg target_arch arm64 '.[] | select((.goarch==$target_arch) and .type=="Archive") | .path' /dist/artifacts.json > /build/plugin-file.txt 
  tar -zxf "/$(cat /build/plugin-file.txt )"
  mv plugin /bin/plugin
EOT

RUN rm -rf /build

CMD ["/bin/plugin"]