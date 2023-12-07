#!/bin/bash
sleep 5
goose postgres "host=$DB_HOST user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME" up