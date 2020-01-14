# web容器化编排服务
gin, mysql, redis搭建docker服务。功能：**gin提供web接口创建用户和查询用户创建个数，redis记录创建用户的个数，用户信息记录到mysql。**

- gin
- mysql
- redis

[源码地址](https://github.com/liangjfblue/docker-all/tree/master/link-gin-db)

## docker的烦恼
一个web架构，往往是涉及多个组件的，通过网络通信来组成一个统一项目来一致对外提供服务，如果是单纯的使用docker，那么久需要我们为每个组件启动都写一个Dockerfile，
并且一个个创建镜像，启动容器。

这样的缺点很明显，在微服务或者组件很多，各个组件服务之间的依赖错综复杂，单单是因为容器的启动顺序，那么就头疼死人了。容器编排正式在这种需求下产生的。

## docker-compose
docker-compose是docker官方的容器编排工具，主要作用是按依赖顺序一次性拉起所有服务的容器。

docker-compose文件是一个定义服务、 网络和卷的 YAML 文件 。Compose 文件的默认路径是 ./docker-compose.yml

## 项目演练
### 1、docker-compose.yaml内容

    version: '3'
    
    services:
    
      web:
        build:
          context: .
          dockerfile: Dockerfile
        labels:
          com.example.description: "docker link web"
          com.example.department: "laingjf"
        ports:
          - "7070:7070"
        deploy:
          replicas: 2
          update_config:
            parallelism: 2
            delay: 10s
          restart_policy:
            condition: on-failure #no  always  on-failure  unless-stopped
        depends_on:
          - redis
          - mysql
        environment:
          REDIS_URL: redis:6379
        networks:
          - backend
    
      redis:
        image: "redis:alpine"
        networks:
          - backend
        restart: on-failure
    
      mysql:
        image: mysql:5.7
        volumes:
          - "./db:/var/lib/mysql"
        restart: on-failure
        environment:
          MYSQL_ROOT_PASSWORD: 123456
          MYSQL_DATABASE: db_docker_link
          MYSQL_USER: root
          MYSQL_PASSWORD: 123456
        networks:
          - backend
    
    networks:
      backend:


### 2、web的Dockerfile文件

    FROM golang:1.13 AS build
    RUN mkdir /build
    WORKDIR	/build
    ADD . .
    WORKDIR	/build/cmd
    ENV GO11MODULE=on
    ENV GOPROXY=https://goproxy.cn,direct
    RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
    
    FROM scratch AS prod
    COPY --from=build /build/cmd/ .
    CMD ["./cmd"]

### 3、目录组织架构

    .
    ├── cmd
    │   ├── cmd
    │   ├── config.yaml
    │   └── main.go
    ├── config
    │   └── config.go
    ├── db
    ├── docker-compose.yaml
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── controllers
    │   │   ├── base
    │   │   │   └── result.go
    │   │   ├── handle.go
    │   │   └── user
    │   │       ├── create.go
    │   │       ├── delete.go
    │   │       ├── get.go
    │   │       ├── init.go
    │   │       ├── login.go
    │   │       ├── logintotal.go
    │   │       └── proto.go
    │   ├── db
    │   │   └── redis
    │   │       ├── client.go
    │   │       └── init.go
    │   ├── models
    │   │   ├── init.go
    │   │   └── tb_user.go
    │   ├── router
    │   │   ├── init.go
    │   │   └── mid
    │   │       └── auth.go
    │   └── server
    │       └── server.go
    ├── pkg
    │   ├── auth
    │   │   └── auth.go
    │   ├── errno
    │   │   └── code.go
    │   └── token
    │       └── token.go
    └── README.md


### 4、执行sudo docker-compose up命令

    Recreating redis ... 
    Starting linkgindb_mysql_1 ... 
    Recreating redis
    Recreating redis ... done
    Recreating linkgindb_web_1 ... 
    Recreating linkgindb_web_1 ... done
    Attaching to linkgindb_mysql_1, linkgindb_redis_1, linkgindb_web_1
    mysql_1  | 2020-01-14 06:04:59+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 5.7.28-1debian9 started.
    mysql_1  | 2020-01-14 06:04:59+00:00 [Note] [Entrypoint]: Switching to dedicated user 'mysql'
    redis_1  | 1:C 14 Jan 2020 06:05:00.387 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
    redis_1  | 1:C 14 Jan 2020 06:05:00.387 # Redis version=5.0.7, bits=64, commit=00000000, modified=0, pid=1, just started
    mysql_1  | 2020-01-14 06:04:59+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 5.7.28-1debian9 started.
    web_1    | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
    redis_1  | 1:C 14 Jan 2020 06:05:00.387 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
    web_1    |  - using env:	export GIN_MODE=release
    web_1    |  - using code:	gin.SetMode(gin.ReleaseMode)
    web_1    | 
    web_1    | [GIN-debug] POST   /v1/user                  --> link-gin-db/internal/controllers/user.(*User).Create-fm (2 handlers)
    web_1    | [GIN-debug] POST   /login                    --> link-gin-db/internal/controllers/user.(*User).Login-fm (2 handlers)
    web_1    | [GIN-debug] GET    /v1/user/:uid             --> link-gin-db/internal/controllers/user.(*User).Get-fm (3 handlers)
    web_1    | [GIN-debug] GET    /v1/func/logintotal       --> link-gin-db/internal/controllers/user.(*User).LoginTotal-fm (2 handlers)
    mysql_1  | 2020-01-14T06:04:59.458178Z 0 [Warning] TIMESTAMP with implicit DEFAULT value is deprecated. Please use --explicit_defaults_for_timestamp server option (see documentation for more details).
    redis_1  | 1:M 14 Jan 2020 06:05:00.388 * Running mode=standalone, port=6379.
    mysql_1  | 2020-01-14T06:04:59.459272Z 0 [Note] mysqld (mysqld 5.7.28) starting as process 1 ...
    redis_1  | 1:M 14 Jan 2020 06:05:00.388 # WARNING: The TCP backlog setting of 511 cannot be enforced because /proc/sys/net/core/somaxconn is set to the lower value of 128.
    redis_1  | 1:M 14 Jan 2020 06:05:00.388 # Server initialized
    redis_1  | 1:M 14 Jan 2020 06:05:00.388 # WARNING overcommit_memory is set to 0! Background save may fail under low memory condition. To fix this issue add 'vm.overcommit_memory = 1' to /etc/sysctl.conf and then reboot or run the command 'sysctl vm.overcommit_memory=1' for this to take effect.
    mysql_1  | 2020-01-14T06:04:59.462262Z 0 [Note] InnoDB: PUNCH HOLE support available
    mysql_1  | 2020-01-14T06:04:59.462279Z 0 [Note] InnoDB: Mutexes and rw_locks use GCC atomic builtins
    redis_1  | 1:M 14 Jan 2020 06:05:00.388 # WARNING you have Transparent Huge Pages (THP) support enabled in your kernel. This will create latency and memory usage issues with Redis. To fix this issue run the command 'echo never > /sys/kernel/mm/transparent_hugepage/enabled' as root, and add it to your /etc/rc.local in order to retain the setting after a reboot. Redis must be restarted after THP is disabled.
    mysql_1  | 2020-01-14T06:04:59.462292Z 0 [Note] InnoDB: Uses event mutexes
    mysql_1  | 2020-01-14T06:04:59.462295Z 0 [Note] InnoDB: GCC builtin __atomic_thread_fence() is used for memory barrier
    mysql_1  | 2020-01-14T06:04:59.462298Z 0 [Note] InnoDB: Compressed tables use zlib 1.2.11
    mysql_1  | 2020-01-14T06:04:59.462300Z 0 [Note] InnoDB: Using Linux native AIO
    redis_1  | 1:M 14 Jan 2020 06:05:00.388 * DB loaded from disk: 0.000 seconds
    redis_1  | 1:M 14 Jan 2020 06:05:00.388 * Ready to accept connections
    mysql_1  | 2020-01-14T06:04:59.462518Z 0 [Note] InnoDB: Number of pools: 1
    mysql_1  | 2020-01-14T06:04:59.462616Z 0 [Note] InnoDB: Using CPU crc32 instructions
    mysql_1  | 2020-01-14T06:04:59.463909Z 0 [Note] InnoDB: Initializing buffer pool, total size = 128M, instances = 1, chunk size = 128M
    mysql_1  | 2020-01-14T06:04:59.472452Z 0 [Note] InnoDB: Completed initialization of buffer pool
    mysql_1  | 2020-01-14T06:04:59.474017Z 0 [Note] InnoDB: If the mysqld execution user is authorized, page cleaner thread priority can be changed. See the man page of setpriority().
    mysql_1  | 2020-01-14T06:04:59.485601Z 0 [Note] InnoDB: Highest supported file format is Barracuda.
    mysql_1  | 2020-01-14T06:04:59.493587Z 0 [Note] InnoDB: Creating shared tablespace for temporary tables
    mysql_1  | 2020-01-14T06:04:59.493658Z 0 [Note] InnoDB: Setting file './ibtmp1' size to 12 MB. Physically writing the file full; Please wait ...
    mysql_1  | 2020-01-14T06:04:59.532810Z 0 [Note] InnoDB: File './ibtmp1' size is now 12 MB.
    mysql_1  | 2020-01-14T06:04:59.533895Z 0 [Note] InnoDB: 96 redo rollback segment(s) found. 96 redo rollback segment(s) are active.
    mysql_1  | 2020-01-14T06:04:59.533913Z 0 [Note] InnoDB: 32 non-redo rollback segment(s) are active.
    mysql_1  | 2020-01-14T06:04:59.534481Z 0 [Note] InnoDB: 5.7.28 started; log sequence number 12441435
    mysql_1  | 2020-01-14T06:04:59.534712Z 0 [Note] InnoDB: Loading buffer pool(s) from /var/lib/mysql/ib_buffer_pool
    mysql_1  | 2020-01-14T06:04:59.534951Z 0 [Note] Plugin 'FEDERATED' is disabled.
    mysql_1  | 2020-01-14T06:04:59.536917Z 0 [Note] InnoDB: Buffer pool(s) load completed at 200114  6:04:59
    mysql_1  | 2020-01-14T06:04:59.542218Z 0 [Note] Found ca.pem, server-cert.pem and server-key.pem in data directory. Trying to enable SSL support using them.
    mysql_1  | 2020-01-14T06:04:59.542239Z 0 [Note] Skipping generation of SSL certificates as certificate files are present in data directory.
    mysql_1  | 2020-01-14T06:04:59.543302Z 0 [Warning] CA certificate ca.pem is self signed.
    mysql_1  | 2020-01-14T06:04:59.543347Z 0 [Note] Skipping generation of RSA key pair as key files are present in data directory.
    mysql_1  | 2020-01-14T06:04:59.543741Z 0 [Note] Server hostname (bind-address): '*'; port: 3306
    mysql_1  | 2020-01-14T06:04:59.543767Z 0 [Note] IPv6 is available.
    mysql_1  | 2020-01-14T06:04:59.543776Z 0 [Note]   - '::' resolves to '::';
    mysql_1  | 2020-01-14T06:04:59.543794Z 0 [Note] Server socket created on IP: '::'.
    mysql_1  | 2020-01-14T06:04:59.546289Z 0 [Warning] Insecure configuration for --pid-file: Location '/var/run/mysqld' in the path is accessible to all OS users. Consider choosing a different directory.
    mysql_1  | 2020-01-14T06:04:59.552780Z 0 [Note] Event Scheduler: Loaded 0 events
    mysql_1  | 2020-01-14T06:04:59.552946Z 0 [Note] mysqld: ready for connections.
    mysql_1  | Version: '5.7.28'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  MySQL Community Server (GPL)
    web_1    | 2020/01/14 06:05:01 server start...


### 5、查看docker-compose up 后的容器启动情况

    liangjf@blue:~/ljf_home/code/go_home/study/docker-all/link-gin-db$ sudo docker container ls
    CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                    NAMES
    8c269603fe64        linkgindb_web       "./cmd"                  3 minutes ago       Up 3 minutes        0.0.0.0:7070->7070/tcp   linkgindb_web_1
    35dfebec112c        mysql:5.7           "docker-entrypoint.s…"   8 minutes ago       Up 6 minutes        3306/tcp, 33060/tcp      linkgindb_mysql_1
    24b535023aeb        redis:alpine        "docker-entrypoint.s…"   15 minutes ago      Up 6 minutes        6379/tcp                 redis
    

### 6、请求web RESTful接口
创建用户:

    [POST] http://172.16.7.16:7070/v1/user
    {
        "user_name":"liangjf",
        "user_pwd":"123456"
    }

获取创建用户数量：

    [GET] http://172.16.7.16:7070/v1/func/logintotal

返回结果

    {
        "code": 1,
        "msg": "ok",
        "data": {
            "total": 1
        }
    }