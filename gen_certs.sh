#!/bin/bash -xe

rm -rf "./test_certs"
mkdir "./test_certs"
pushd "./test_certs"

openssl req \
    -newkey rsa:2048 \
    -nodes \
    -days 3650 \
    -x509 \
    -keyout ca.key \
    -out ca.crt \
    -subj "/CN=*"

openssl req \
    -newkey rsa:2048 \
    -nodes \
    -keyout server.key \
    -out server.csr \
    -subj "/C=US/ST=California/L=San Jose/O=cubed/OU=TestServer/CN=*"

openssl x509 \
    -req \
    -days 3650 \
    -sha256 \
    -in server.csr \
    -CA ca.crt \
    -CAkey ca.key \
    -CAcreateserial \
    -out server.crt \
    -extfile <(echo subjectAltName = IP:127.0.0.1)

openssl req \
    -newkey rsa:2048 \
    -nodes \
    -keyout client.key \
    -out client.csr \
    -subj "/C=US/ST=California/L=San Jose/O=cubed/OU=TestClient/CN=*"

openssl x509 \
    -req \
    -days 3650 \
    -sha256 \
    -in client.csr \
    -CA ca.crt \
    -CAkey ca.key \
    -CAcreateserial \
    -out client.crt

openssl x509 \
    -req \
    -days 3650 \
    -sha256 \
    -in client.csr \
    -CA ca.crt \
    -CAkey ca.key \
    -CAcreateserial \
    -out client-spiffe.crt \
    -extfile <(echo subjectAltName = URI:spiffe://test/client)
