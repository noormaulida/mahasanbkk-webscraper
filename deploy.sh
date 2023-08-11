#!/bin/sh
GOOS=linux GOARCH=amd64 go build cmd/main.go
sudo systemctl daemon-reload
sudo service mahasan-app stop
sudo service mahasan-app start
sudo service nginx restart