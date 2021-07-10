#!/bin/bash

# 初回のみClone、以降はPullする
if cd TEAMC; then
  git pull;
else
  git clone $1 TEAMC;
  cd TEAMC
fi

docker-compose down -v
docker-compose build
docker-compose up -d