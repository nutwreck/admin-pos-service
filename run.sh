#!/bin/bash

# Clear terminal
clear

# Generate Docs
$HOME/go/bin/swag init

# Run the application
go run main.go