#!/bin/bash
cd /home/docker/sicerdas-indonesia-server
docker compose exec go /app/utils/skoring/main
docker compose exec go /app/utils/publish/main
