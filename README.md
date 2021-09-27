# FriendlyAlmond Backend

## 项目目录结果说明
├── bin     编译后可执行文件目录   
├── cmd     所有可执行文件的 main 入口    
├── conf    配置文件   
├── controller   api 的 controller，  
├── deployment   DockerFile，docker-compose和docker 镜像相关配置    
├── go.mod   
├── go.sum   
├── log          本地运行的log，已加入 git ignore   
├── Makefile     项目主要的 makefile，所有编译相关的命令都在里面。   
├── pkg          项目公用的代码库   
├── proto        pb文件，包含 proto 文件和编译后的go文件    
├── README    
├── scripts      一些需要用到的脚本    
└── service      所有 rpc handler 相关的业务逻辑，类似 controller 的作用

## 配置文件说明
本地配置文件统一连到坐标悉尼的测试服 `test.opsnft.net` 这包括了 mongodb 和 redis，各位本地测试的话请自行修改 `conf` 目录下**没有** `_online`后缀的配置文件，修改后的本地测试配置请不要提交到 git 上。

### 本地调试方法
请先安装 `consul` 并启动服务……

```
brew install consul
brew services start consul
```
然后执行 `make run api` 启动 Web 服务，否则其他服务不能收到请求。

最后使用类似 `make run <service>` 的命令进行进行本地调试，其中 `<service>` 就是要调试启动的模块，比如 `contract`、`oop` 等。

相关服务的配置文件是 `conf/<service>.yaml`，带有 `_online` 的配置文件用于线上 docker 读取，如果要修改配置文件，请不要 push 本地配置，如果改了线上配置，记得也提交。

### 生成 API 文档
`make swag`


## 编译
### 编译单个服务
```bash
make property
```
cmd 下的文件夹名字对应的服务名，可执行文件名，make 的依赖名。详情见makefile

### 编译所有可执行文件
`make build`

### 打包
```bash
make docker_build
```
会重新编译所有项目，并把 conf 和 bin 目录复制到镜像里面去。详情见 `deployment/dockerfile` 。

## 运行
### 运行单个服务
```bash
./bin/api -f conf/api.api   # 注意：所有服务都要显示的指定配置文件的路径
```

### 运行所有服务
运行所有服务前请提前启动以下依赖组件，并暴露好相关端口。如需改变这些组件的端口，请在 conf 相关的配置也相应改变。
consule 8500
mongo  27017
redis  6379

启动好以上服务后，再通过 docker-compose 启动所有其他服务，你也可以停掉其中一个服务，并手动从 bin 目录启动一个服务。
```bash
make up  # docker-compose 配置在 deployment/docker-compose-local.yaml
```
以上命令会运行除 api 服务以外所有的服务。由于 api 需要暴露端口，走 host 网络模式无法直接运行。需单独运行。

停止所有服务
```bash
make down
```

## 查看线上服务运行状态
通常来说，我们都可以在  [http://test.opsnft.net:8500/ui/whoops/services](http://test.opsnft.net:8500/ui/whoops/services) 查看服务运行状态，当然这并不一定保证服务运行正确，但至少如果报错了那服务肯定是挂了的。
