#! /bin/bash

cd ../dodson-site
ng build
cd ../server
go run server.go
