#!/bin/sh

openssl genrsa -des3 -passout "pass:secret" -out myCA.key 2048 

openssl req -x509  -passin "pass:secret" -new -nodes -key myCA.key -sha256 -days 1825 -subj "/C=US/ST=State/L=City/O=Company Inc./OU=IT/CN=localhost" -out myCA.pem
