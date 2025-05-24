#!/bin/bash

command="$1"

if [ -z "$command" ]; then
    command="start"
fi

ProjectRoot="$(dirname "$(dirname "$0")")"
export AMBULANCE_API_ENVIRONMENT="Development"
export AMBULANCE_API_PORT="8080"
export AMBULANCE_API_MONGODB_USERNAME="root"
export AMBULANCE_API_MONGODB_PASSWORD="neUhaDnes"

function mongo {
    docker compose --file "${ProjectRoot}/deployments/docker-compose/compose.yaml" $@
}

case "$command" in
    "openapi")
        docker run --rm -ti -v "${ProjectRoot}:/local" openapitools/openapi-generator-cli generate -c /local/scripts/generator-cfg.yaml
        ;;
    "start")
        try_mongo_down=false
        trap 'try_mongo_down=true' EXIT
        
        mongo up --detach
        go run "${ProjectRoot}/cmd/ambulance-api-service"
        
        if [ "$try_mongo_down" = true ]; then
            mongo down
        fi
        ;;
    "test")
        go test -v ./...
        ;;
    "docker")
        docker build -t patriksabol/ambulance-wl-webapi:local-build -f ${ProjectRoot}/build/docker/Dockerfile .
        ;;
    "mongo")
        mongo up
        ;;
    *)
        echo "Unknown command: $command" >&2
        exit 1
        ;;
esac