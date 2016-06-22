FROM opower/opentsdb

RUN cd /opt/opentsdb && git fetch origin v2.2.0:2.2.0 && git checkout 2.2.0
RUN cd /opt/opentsdb && bash ./build.sh
ADD opentsdb.conf /etc/opentsdb.conf
ADD start_opentsdb.sh /opt/tsdb-docker/start_opentsdb.sh
