# Microservice Template #
Microservices using [ZMQ](http://zeromq.org/) and Go.
You need docker and docker compose to run this.

Pull requests welcome :smiley:

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
You should see something like this:
```
worker_1  | Received request [Hello 9]
client_1  | Client 0 Received [Hello 9]
client_1  | It took: 10.074335 seconds.
```
It shows that it took more or less 10 seconds to process all the requests for 1 client. This example client is doing 10 requests, each request takes 1 second to process on the worker for this demo.

You can define the number of clients:
```
docker-compose up --scale client=2 
```

stop:
```
docker-compose down
```
## Scaling it up ##

With more clients the time to process everyone increases, so we need to make it scalable, you can also define the number of threads for each worker in the docker-compose.yaml file, or run any number of workers with:
```
docker-compose up --scale client=2 --scale worker=2
```

## Architecture ##
### Router ###

The router container will receive all the requests from the clients, his job is to do a fairly distribution of the workload.  
Is based on the request-reply broker of the ZMQ guide: http://zguide.zeromq.org/page:all#Shared-Queue-DEALER-and-ROUTER-sockets

### Worker ###

The worker connects to a router and is always waiting for requests, its has a internal router to distribute the work by his threads. If only one worker process is preferred the router can then be removed making a direct communication between the clients and the worker.

### Logger ###
The logger container work as a log sink, it pulls the logs from the workers and write to a log file.  
It uses: https://github.com/natefinch/lumberjack for log rotation.  
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
