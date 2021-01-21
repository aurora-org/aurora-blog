FROM alpine:latest
WORKDIR /app
COPY . .
EXPOSE 8080 8888
ENTRYPOINT ["nohup", "./aurora", "-conf", "./config/runtime.config"]
