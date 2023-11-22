#!/usr/bin/env bash 
go build -o /app/utils/publish/main /app/utils/publish/main.go
echo "Berhasil Build Publish  main.go for cronjob"

go build -o /app/utils/skoring/main /app/utils/skoring/main.go
echo "Berhasil Build Skoring main.go for cronjob"
cd /app
air