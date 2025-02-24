#!/bin/bash

mc alias set minio http://minio-container:9000 $MINIO_ROOT_USER $MINIO_ROOT_PASSWORD
mc mb --ignore-existing minio/"$VIDEO_BUCKET"
mc anonymous set public minio/"$VIDEO_BUCKET"
mc mb --ignore-existing minio/"$IMAGE_BUCKET"
mc anonymous set public minio/"$IMAGE_BUCKET"
mc mb --ignore-existing minio/"$DOCUMENT_BUCKET"
mc anonymous set public minio/"$DOCUMENT_BUCKET"

for file in data/*.jpg; do
    if [ -f "$file" ]; then
        mc put "$file" minio/"$IMAGE_BUCKET"
    fi
done

for file in data/*.mp4; do
    if [ -f "$file" ]; then
        mc put "$file" minio/"$VIDEO_BUCKET"
    fi
done

mc put data/1.pdf minio/"$DOCUMENT_BUCKET"/1
mc put data/2.pdf minio/"$DOCUMENT_BUCKET"/2
mc put data/3.pdf minio/"$DOCUMENT_BUCKET"/3