# Monitoring Job

## Purpose

Build an executable to read the Temp & Humidity from a DHT22 and push it over to an API.

## Build it
$ docker build . -t mj_compiler
$ docker run -v $(pwd):/app mj_compiler

The executable will be: ./app