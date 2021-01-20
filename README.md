## Aurora Blog System

Blog System Based on Golang named Aurora

### Frameworks

- iris
- gorm
- viper
- graphql-go

### Pre-environment

- Docker
- Golang
- GUN Make

### Steps

1. `git clone git@github.com:qmdx00/aurora-blog.git`
2. config local go environment via below code.
    ```shell
    export GOPROXY=goproxy.cn
    export GO111MODULE=on
    ```
3. cd project root directory then run `go mod tidy` to install dependencies.
4. type command `docker run -d --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD={your-password} mysql:latest` to start a
   mysql container.
5. run sql in mysql
   container `SET GLOBAL sql_mode = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';`
6. config `aurora-blog/config/runtime.config.json` file to update database config.
7. type command `make config` to generate database table.
8. type command `make dev` to start project.

> You can visit http://127.0.0.1:8080 for Restful API. </br>
> And visit http://127.0.0.1:8888/graphql for GraphQL server.
