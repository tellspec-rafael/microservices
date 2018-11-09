# Microservices #
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

## Logging ##
Logger container work as a log sink, it pulls the logs from the workers and write to a rotative log file.  
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
