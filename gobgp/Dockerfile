FROM chenchun/quagga

ENV HOME /root
WORKDIR /root
RUN go get -u github.com/golang/dep/cmd/dep
RUN mkdir -p /go/src/github.com/osrg \
        && git clone https://github.com/osrg/gobgp.git /go/src/github.com/osrg/gobgp \
        && cd /go/src/github.com/osrg && dep ensure \
        && go install github.com/osrg/gobgp/cmd/gobgp \
        && go install github.com/osrg/gobgp/cmd/gobgpd
