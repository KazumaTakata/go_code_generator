#!/bin/bash

go run generator.go >sample.go
go fmt sample.go
echo "generated code"
echo "--------------"
cat sample.go
