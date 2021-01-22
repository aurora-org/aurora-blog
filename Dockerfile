FROM alpine:latest
WORKDIR /app
ENV ZONEINFO=/app/zoneinfo.zip
COPY . .
EXPOSE 8080 8888
ENTRYPOINT ["nohup", "./aurora", "-conf", "./config/deploy.config"]
