FROM golang:1.11-stretch

RUN apt-get update \
    && apt-get install -qy --no-install-recommends \
    supervisor quagga telnet tcpdump vim net-tools \
    && rm -rf /var/lib/apt/lists/*
ADD supervisord.conf /etc/supervisor/conf.d/supervisord.conf
RUN sed -i 's/zebra=no/zebra=yes/g' /etc/quagga/daemons \
    && cp /usr/share/doc/quagga/examples/zebra.conf.sample /etc/quagga/zebra.conf

CMD ["/usr/bin/supervisord"]
