#!/bin/sh
GOOS=linux GOARCH=amd64 go build cmd/main.go
sudo service mahasan-app restart
sudo service nginx restart