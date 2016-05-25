#!/bin/bash

echo "Sleeping for 15 seconds to give HBase time to warm up"
sleep 15

CHECK="$(echo "exists 'tsdb'" | /opt/hbase/bin/hbase shell | grep -oh "Table tsdb does not exist")"
while [ "$CHECK" == "Table tsdb does not exist" ]
do
    echo "Attemping to create OpenTSDB tables"
    /opt/opentsdb/src/create_table.sh
    sleep 5
    CHECK="$(echo "exists 'tsdb'" | /opt/hbase/bin/hbase shell | grep -oh "Table tsdb does not exist")"
done

echo "Starting tcollector"
#/opt/tcollector/startstop start &

echo "Starting OpenTSDB"
/opt/opentsdb/build/tsdb tsd --port=4242 --staticroot=/opt/opentsdb/build/staticroot --cachedir=/tmp --auto-metric
