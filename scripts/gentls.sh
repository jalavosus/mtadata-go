#!/bin/bash

mkdir tlscerts
cd tlscerts

openssl ecparam -name secp521r1 -genkey -noout -out ca.key
openssl ecparam -name secp521r1 -genkey -noout -out server.key

openssl req -new -x509 -key ca.key -sha256 -subj "/C=US/ST=NY/O=Carl" -days 365 -out ca.cert
openssl req -x509 -days 365 -CA ca.cert -CAkey ca.key -key server.key -out server.cert -config ../scripts/gentls.conf -extensions 'v3_req'