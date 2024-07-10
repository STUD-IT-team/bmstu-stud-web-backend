#!/bin/bash

mc config host add minio http://minio:9000 $MINIO_ROOT_USER $MINIO_ROOT_PASSWORD
mc mb --ignore-existing minio/"$VIDEO_BUCKET"
mc anonymous set public minio/"$VIDEO_BUCKET"
mc mb --ignore-existing minio/"$IMAGE_BUCKET"
mc anonymous set public minio/"$IMAGE_BUCKET"
mc mb --ignore-existing minio/"$DOCUMENT_BUCKET"
mc anonymous set public minio/"$DOCUMENT_BUCKET"
mc put data/main_vid.mp4 minio/"$VIDEO_BUCKET"
mc put data/arch.png minio/"$IMAGE_BUCKET"