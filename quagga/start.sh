#! /bin/bash
mkdir -p /run/quagga
chmod 777 /run/quagga
/usr/bin/supervisord
