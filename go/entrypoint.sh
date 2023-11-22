#!/usr/bin/env bash 
go build -o /app/utils/skoring/main /app/utils/skoring/main.go
echo "Berhasil Build Skoring main.go for cronjob"
cd /app
air