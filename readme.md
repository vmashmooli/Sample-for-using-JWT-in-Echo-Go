# Simple_JWT_in_Golang

Simple web API example for using JWT in GoLang

## Description

Creating a simple web service by Echo and APIs to create and check the token, and Using JSON file to keep configuation.
Items given in this example:

1 — How to create and loading configuration from JSON files
2 — How to run a simple web service by Echo
3 — How to Routing APIs in Echo
4 — How to create a token by user id and user name
5 — How to check received token validity 

## Getting Started

Step 1 — Creating a JSON file for configuration
Step 2 — Reading the config.json file
Step 3 — Creating a web service by Echo 
Step 4 — Creating a token for the user
Step 5 — Checking the received token validity

### Dependencies

* ubuntu 16 or higher
* Golang 1.11 or higher
* Echo 4.7.0 or higher
* import JWT and fmt, log, os, time, errors, encoding/json, net/http packages

### Executing program

* go mod init
* go mod tidy
* go run main.go / go run .
* http://localhost:1707/api/v1/create_token
* http://localhost:1707/api/v1/check_token?token={token value}

### Extra links

Echo Go guide:  https://echo.labstack.com/guide
JWT Debugger: https://jwt.io
Secure Password Generator: https://passwordsgenerator.net

## Authors

vahid mashmooli
vahidmashmooli@gmail.com

## Version History

* 0.001