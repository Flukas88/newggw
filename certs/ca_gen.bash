#!/bin/sh

openssl genrsa -des3 -passout "pass:secret" -out ca.key 2048 

openssl req -x509  -passin "pass:secret" -new -nodes -key ca.key -sha256 -days 1825 -subj "/C=US/ST=State/L=City/O=Company Inc./OU=IT/CN=localhost" -out ca.pem
ln -s ca.pem ca.cert
