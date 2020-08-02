#!/bin/sh
FILES="../../*.csv"
for i in $FILES
do
    go run main.go $i
done
