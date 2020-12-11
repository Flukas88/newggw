#!/bin/sh

if [ "$#" -ne 1 ]
then
  echo "Usage: Must supply a domain"
  exit 1
fi

DOMAIN=$1


openssl genrsa -out $DOMAIN.key 2048
openssl req -new -subj "/C=AU/ST=NSW/L=Melbourn/O=CoupaInvoiceSmash/OU=Development/CN=${DOMAIN}/emailAddress=test@test.com" -key $DOMAIN.key -out $DOMAIN.csr



cat > $DOMAIN.ext << EOF
[ req ]
prompt = no
distinguished_name = server_distinguished_name
req_extensions = v3_req

[ server_distinguished_name ]
commonName = $DOMAIN
stateOrProvinceName = NSW
countryName = AU
emailAddress = test@test.com
organizationName = Coupa InvoiceSmash
organizationalUnitName = Development

[ v3_req ]
basicConstraints = CA:FALSE
keyUsage = nonRepudiation, digitalSignature, keyEncipherment
subjectAltName = @alt_names

[ alt_names ]
DNS.0 = $DOMAIN
EOF

openssl x509 -req  -passin pass:secret -in $DOMAIN.csr -CA ca.pem -CAkey ca.key -CAcreateserial -out $DOMAIN.crt -days 825 -sha256 -extfile $DOMAIN.ext