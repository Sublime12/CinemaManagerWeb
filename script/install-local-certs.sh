#!/bin/bash
set -e

mkdir -p certs

mkcert -install

mkcert \
  -cert-file certs/cinema.localhost.pem \
  -key-file certs/cinema.localhost-key.pem \
  "cinema.localhost" \
  "*.cinema.localhost"
