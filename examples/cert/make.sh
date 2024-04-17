#!/bin/sh
openssl req -new -newkey rsa:4096 -x509 -sha256 -days 36500 -nodes -out cert.crt -keyout key.key
