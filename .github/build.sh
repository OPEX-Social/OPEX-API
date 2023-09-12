#!/bin/bash

# Navigate to the directory where Go code is located
cd ./API

# Remove the previous executable
rm -rf ./../OPEX-API.exe

# Build the Go project into an executable
go build -o ./../OPEX-API.exe

# Pause the script to view any errors
read -p "Press any key to close this window..."
