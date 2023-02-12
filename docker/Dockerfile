FROM alpine

LABEL version="v0.6.1" \
    maintainer="wang@rehiy.com"

ADD tdp-cloud.gz /usr/local/bin/tdp-cloud

RUN chmod +x /usr/local/bin/tdp-cloud \
    && mkdir -p /var/lib/tdp-cloud \
    && mkdir -p /var/log/tdp-cloud \
    && mkdir -p /etc/tdp-cloud

ENTRYPOINT ["tdp-cloud", "server", "-c", "/etc/tdp-cloud/config.yml"]

EXPOSE 7800
