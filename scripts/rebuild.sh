#!/bin/bash

imageName=xx:__image__
containerName=__container__

docker build -t $imageName -f Dockerfile .

echo Delete old container...
docker rm -f $containerName

echo Run new container...
docker run -d -p 1313:1313 --network="proxy" --name $containerName $imageName

echo done
