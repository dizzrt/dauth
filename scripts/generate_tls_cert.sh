#!/bin/sh

# exit when any command fails
set -e

TLS_DIR="./configs/tls"
CERT_FILE="$TLS_DIR/server.crt"
KEY_FILE="$TLS_DIR/server.key"
VALID_DAYS=365
COMMON_NAME="localhost"
SAN_LIST="DNS:localhost,DNS:127.0.0.1,IP:0.0.0.0,IP:127.0.0.1"

mkdir -p "$TLS_DIR"

echo "Starting to generate development TLS certificate..."
openssl req -x509 -newkey rsa:4096 \
    -keyout "${KEY_FILE}" \
    -out "${CERT_FILE}" \
    -days "${VALID_DAYS}" \
    -nodes \
    -subj "/CN=${COMMON_NAME}" \
    -addext "subjectAltName=${SAN_LIST}"

chmod 600 "$KEY_FILE"
chmod 644 "$CERT_FILE"

echo "Certificate generation completed:"
echo "  Certificate file: $CERT_FILE"
echo "  Private key file: $KEY_FILE"
echo "  Validity period: $VALID_DAYS days"
echo "  Applicable domains/IPs: $(echo "$SAN_LIST" | sed 's/,/ | /g')"
