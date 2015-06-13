#!/bin/bash

cd public/app;
grunt clean && grunt build;
cd ../dist;
go-bindata -o=../../static.go ./...;
cd ../../;
go build ./...;
echo "Leggo! :-)";