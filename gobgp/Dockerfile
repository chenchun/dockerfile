FROM chenchun/quagga

ENV HOME /root
WORKDIR /root
RUN go get github.com/osrg/gobgp/cmd/gobgp && go get github.com/osrg/gobgp/cmd/gobgpd && ln -s /go/src/github.com/osrg/gobgp .
