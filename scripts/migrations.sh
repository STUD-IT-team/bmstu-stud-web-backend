#!/bin/bash
. ../deployments/db.env
cd ..
while getopts o: flag
do
    case "${flag}" in
        o) option=${OPTARG}
           goose -dir=./migrations postgres "host=$DBHOST_MAKE user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB" $option
           exit;;
    esac
done