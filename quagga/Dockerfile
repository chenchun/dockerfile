FROM golang:1.11-stretch

RUN apt-get update \
    && apt-get install -qy --no-install-recommends \
    supervisor quagga telnet tcpdump vim net-tools wget lv \
    && rm -rf /var/lib/apt/lists/*
ADD supervisord.conf /etc/supervisor/conf.d/supervisord.conf
ADD daemons /etc/quagga/daemons
ADD start.sh /opt
RUN sed -i 's/zebra=no/zebra=yes/g' /etc/quagga/daemons \
    && cp /usr/share/doc/quagga-core/examples/zebra.conf.sample /etc/quagga/zebra.conf

CMD ["/opt/start.sh"]
