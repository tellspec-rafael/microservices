# Microservice Template #
Microservices using ZMQ and Go.
You need docker and docker compose to run this.

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

## Architecture ##
### Router ###

The router container will receive all the requests from the clients, his job is to do a fairly distribution of the workload.  
Is based on the request-reply broker of the ZMQ guide: http://zguide.zeromq.org/page:all#Shared-Queue-DEALER-and-ROUTER-sockets

### Worker ###

The worker connects to a router and is always waiting for requests, its has a internal router to distribute the work by his worker threads. If only one worker process is preferred the router can then be removed making a direct communication between the clients and that one worker.

### Logger ###
The logger container work as a log sink, it pulls the logs from the workers and write to a log file.  
It uses: https://github.com/natefinch/lumberjack for rotative logging.  
Logging options:
```
log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/log.log",
		MaxSize:    2, // megabytes
		MaxBackups: 10,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})
```
