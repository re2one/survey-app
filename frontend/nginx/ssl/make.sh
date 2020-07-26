#!/bin/bash

echo "CA\n";
openssl req -new -x509 -newkey rsa:2048 -keyout cakey.pem -out cacert.pem -days 3650;
echo "Generate key\n";
openssl genrsa -out serverkey.pem -aes128 2048 -days 3650;
openssl rsa -in serverkey.pem -out serverkey.pem;

openssl req -new -key serverkey.pem -out req.pem -nodes;
echo "Sign key\n";
openssl ca -in req.pem -notext -out servercert.pem;
