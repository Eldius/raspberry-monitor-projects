#!/bin/bash

function error {
    echo $1
    exit 1
}


go clean
GOOS=linux GOARCH=arm GOARM=5 go build || error "Error building project"

sftp pi@192.168.100.14 << EOSFTP
    put raspberry-simplemqtt-listener
EOSFTP

retVal=$?
if [ $retVal -ne 0 ]; then
    error "Error sending binary to pi"
fi


ssh pi@192.168.100.14 << EOSSH
    ./raspberry-simplemqtt-listener
    rm ./raspberry-simplemqtt-listener
EOSSH

retVal=$?
if [ $retVal -ne 0 ]; then
    error "Error executing binary on pi"
fi
