#!/usr/bin/env sh



if [ "$#" -gt 0 ]; then
    BINARY=portainerImporter
fi

go build -o /opt/bin/${BINARY} .
