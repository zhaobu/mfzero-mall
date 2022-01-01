#!/bin/bash

allsvc=(admin-api oms pay pms sms sys ums)
if [ $# -eq 0 ]; then
    echo "参数太少 ./docker-compose.sh type servername"
elif [ $1 = 'restart' ]; then
    if [ $# -eq 1 ]; then
        for srv in ${allsvc[@]}; do
            docker-compose -f "deploy/docker-compose.yml" restart $srv
            echo "docker-compose restart $srv"
        done
    else
        docker-compose -f "deploy/docker-compose.yml" restart $2
        echo "docker-compose restart $2"
    fi
elif [ $1 = 'rebuild' ]; then
    docker-compose -f "deploy/docker-compose.yml" up -d --build $2
    # docker-compose -f "deploy/docker-compose.yml" build $2
    # docker-compose -f "deploy/docker-compose.yml" create $2
    # docker-compose -f "deploy/docker-compose.yml" start $2
else
    echo "nothing"
fi
