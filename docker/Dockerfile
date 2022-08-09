#syntax=docker/dockerfile:1.3-labs

FROM alpine
LABEL org.opencontainers.image.source https://github.com/kameshsampath/drone-quickstart
LABEL org.opencontainers.image.authors="Kamesh Sampath<kamesh.sampath@hotmail.com>"

ARG TARGETARCH

RUN apk -Uuv add curl bash jq

ADD ./dist /build

RUN <<EOT
    jq -r --arg target_arch $TARGETARCH  '.[] | select(.goarch==$target_arch) | .path' /build/dist/artifacts.json > /build/pluginfile.txt
    cp $(cat /build/pluginfile.txt) /bin/plugin
    chmod +x /bin/plugin
EOT

RUN rm -rf /build

CMD ["/bin/plugin"]