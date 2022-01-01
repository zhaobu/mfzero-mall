#!/bin/bash
output=bin/server
function build() {
    pwd
    # make clean
    make build
    chmod u+x $output
    return 0
}

if [ $# -ne 2 ]; then
    echo "参数太少 ./start servertype servername"
elif [ $1 = 'api' ]; then
    cd api/$2 && build
    $output -f etc/$2-api.yaml
elif [ $1 = 'service' ]; then
    cd app/$2/service && build
    $output -f etc/$2.yaml
else
    echo "./start api $servername || ./start service $servername"
fi
