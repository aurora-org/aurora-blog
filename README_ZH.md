## Aurora 博客系统

基于 Golang 的博客系统

### 涉及框架

- iris
- gorm
- viper
- graphql-go

### 前置环境

- Docker
- Golang
- GUN Make

### 使用步骤

1. clone该仓库，`git clone git@github.com:qmdx00/aurora-blog.git`
2. 修改本地 Go 环境配置
    ```shell
    export GOPROXY=goproxy.cn
    export GO111MODULE=on
    ```
3. 进入项目根目录，运行 `go mod tidy` 来安装项目依赖.
4. 启动 MySQL Docker 容器 `docker run -d --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD={your-password} mysql:latest`.
5. 修改 MySQL 配置 `SET GLOBAL sql_mode = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';`
6. 修改配置文件中的数据库配置，文件位于 `aurora-blog/config/runtime.config.json`.
7. 项目根目录运行 `make config` 来生成数据表.
8. 运行 `make dev` 来启动项目.

> 通过 http://127.0.0.1:8080 来访问 Restful API. </br>
> 通过 http://127.0.0.1:8888/graphql 访问 GraphQL 服务.
