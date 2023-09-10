#!/bin/bash
docker-compose up
docker exec -it kafka kafka-topics.sh --create \
  --topic event.completed \
  --bootstrap-server localhost:9092 \

make migrate-up
make tidy 
make run