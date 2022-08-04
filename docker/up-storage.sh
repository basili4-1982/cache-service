#!/usr/bin/env sh
docker run  -p 11211:11211  -d memcached memcached -m 64
