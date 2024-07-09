#!/bin/bash

mc config host add minio http://minio:9000 $MINIO_ROOT_USER $MINIO_ROOT_PASSWORD
mc mb --ignore-existing minio/videos
mc anonymous set public minio/videos
mc mb --ignore-existing minio/images
mc anonymous set public minio/images
mc put data/main_vid.mp4 minio/videos
mc put data/arch.png minio/images