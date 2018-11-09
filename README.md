#Microservices
=======
Microservices using ZMQ and Go.
You need docker and docker compose to run this example.

##Build
Build the images running:
```
docker-compose build
```

##Running
Then run the images with:
```
docker-compose up
```

stop:
```
docker-compose down
```
##Scaling up
Run any number or workers:
```
docker-compose up --scale worker=4
```
You can also define the number of threads in each worker in the docker-compose.yaml file.



