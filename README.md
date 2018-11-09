# Microservices #
Microservices using ZMQ and Go.
You need docker and docker compose to run this example.

## Build ##
Build the images:
```
docker-compose build
```

## Running ##
Run the containers with:
```
docker-compose up
```
stop:
```
docker-compose down
```
## Scaling up ##
Run any number or workers:
```
docker-compose up --scale worker=4
```
You can also define the number of threads in each worker in the docker-compose.yaml file.



