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

### Local debug method
Firstly, please make sure you already install the `consul` service and run it……

```
brew install consul
brew services start consul
```
Then use `make run api` run the web service in the docker。

### Generate the API document
`make swag`


## Compile
### Compile a single service
```bash
make login
```

### Compile all service
`make build`

### make docker image
```bash
make docker_build
```
This command will compile all service，and copy the bin file and config file into the image. More detail please look the `deployment/dockerfile` 。

## running
### run a single service
```bash
./bin/api -f conf/api.api   # notes：all of the service need to specify the config file
```
