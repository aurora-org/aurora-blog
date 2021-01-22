FROM alpine:latest
WORKDIR /app
ENV ZONEINFO=/app/zoneinfo.zip
COPY . .
EXPOSE 8080 8888
ENTRYPOINT ["sh", "./entrypoint.sh"]
