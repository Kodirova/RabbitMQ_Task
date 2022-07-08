# Rabbit MQ

Simple overview of use/purpose.

## Description

*  Implementation of pub/sub, Ack/Nack, routing of Rabbit MQ. Work with three services. 
*  First service is for storing phone and selecting through id using REST api
*  Second service is for consumming message throgh queue and exchange logger. Consumers differs from each other through different levels of logging(error, info, debug)
*  Third servise is for implementing producer of message which get response from rest and sent to second service 

## Getting Started


### Installation


1.  Clone the repo
   ```sh
   git clone https://github.com/your_username_/Project-Name.git
   ```
2. Install migrate packages
```bash
$ curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz | tar xvz
```

### Executing program

1. Run docker-compose 
   ```js
  docker-compose up
   ```
2. Create database rabbitmq_task through docker
   ```bash
  create database rabbitmq_task
   ```
3. Migrate sql dependencies
   ```bash
   make migrate-local-up
   ```
4. Runs Rest server
   ```bash
   go run storage/main.go
   ```
5. Runs Consumer
   ```bash
   go run logger/receive_log.go
   ```
6. Runs Publisher
   ```bash
   go run info/emit_log_direct.go
   ```
   

## Help

Any advise for common problems or issues.


## Authors

Kimmy   
https://www.linkedin.com/in/kimmydev/

## Version History

* 0.1
    * Various bug fixes and optimizations
    * See [commit change]() 

## Acknowledgments

Inspiration, code snippets, etc.
*  https://github.com/rabbitmq/rabbitmq-tutorials/blob/master/go
