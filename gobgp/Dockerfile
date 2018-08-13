FROM chenchun/quagga

RUN apt-get update && apt-get install -qy --no-install-recommends wget lv tcpdump && rm -rf /var/lib/apt/lists/*
ENV HOME /root
WORKDIR /root
RUN go get github.com/osrg/gobgp/cmd/gobgp && go get github.com/osrg/gobgp/cmd/gobgpd && ln -s /go/src/github.com/osrg/gobgp .
