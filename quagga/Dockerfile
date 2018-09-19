FROM golang:1.11-stretch

RUN apt-get update \
    && apt-get install -qy --no-install-recommends \
    supervisor quagga telnet tcpdump vim net-tools wget lv \
    && rm -rf /var/lib/apt/lists/*
ADD supervisord.conf /etc/supervisor/conf.d/supervisord.conf
ADD zebra.conf /etc/quagga/zebra.conf
ADD daemons /etc/quagga/daemons

CMD ["/usr/bin/supervisord"]
