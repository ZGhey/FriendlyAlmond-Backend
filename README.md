# FriendlyAlmond Backend

## The Content of the project 
* bin 
  * After compile the project, the executable bin file will in this document.
* cmd
  * The main file of the project.   
* conf
  * Configuration file.
* controller
  * The controller of the API.
* deploy
  * The deployment file including dockerfile, docker-compose file and mysql.
* go.mod   
* go.sum   
* log             
* Makefile
  * Store all the compile command.  
* pkg
  * The public tools package   
* proto 
  * pb file    
* README    
* scripts
  * shell or something script    
* service
  * The RPC handler same like controller

## Local debug mode
Please make sure you already install all the tools, when running local debug mode.

### Environment
* [Go 1.16.3](https://golang.org/dl/)
* Consul 1.10.1
* MySQL 5.7
* [Redis 6.2.3](https://hub.docker.com/_/redis?tab=description)
* About the details of framework pleas look at the go.mod file

### Install Consul
Firstly, installing the `consul` service and run it...

```
brew install consul
brew services start consul
```
Then use `make run api` run the web service in the docker。

Notes: About how to install Consul in docker, please view https://hub.docker.com/_/consul

### Install MySQL in Docker
Please run the shell in deploy/mysql.sh, it will download the MySQL image from official docker hub and install it.

```bash
sh deploy/mysql.sh
```
## Generate the API document
Using the following commend at root path of the project to generate the swagger API for Front-end
```
make swag
```
## running
### run a single service in debug mode
```bash
# Notes：All the service need to specify the config file.
# Notes: If you use local database to run the service, don't forget to modify the config info like MySQL link and Redis link in the config file.
go run cmd/api/main.go -f conf/api.yaml
go run cmd/configuration/main.go -f conf/configuration.yaml
go run cmd/jobModule/main.go -f conf/job_module.yaml
go run cmd/login/main.go -f conf/login.yaml
go run cmd/order/main.go -f conf/order.yaml
```
## Running on server

### Environment
* [Docker 4.0.1](https://docs.docker.com/desktop/mac/release-notes/)
* [Portainer.io 2.1.1](https://docs.portainer.io/v/ce-2.9/start/intro)
* [Consul 1.10.1](https://hub.docker.com/_/consul)
* [Nginx 1.20.1](https://hub.docker.com/_/nginx)
* MySQL 5.7
* About the details of framework pleas look at the go.mod file

## Compile and Running
### Compile a single service
```bash
make api
```
### run a single service
```bash
./bin/api -f conf/api.yaml   # notes：all of the service need to specify the config file
```

### Compile all service
It will create the bin file into cmd document
```
make build
```

### make docker image
```bash
make docker_build
```
This command will compile all service，and copy the bin file and config file into the image. More detail please look the `deploy/Dockerfile` .