FROM alpine:latest

# 设置时区
ENV TZ=Asia/Shanghai
#RUN apk update && apk add --no-cache tzdata openntpd \
#    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone ;

WORKDIR /root/

COPY ./server ./


# 暴露应用的端口
EXPOSE 9999

# 设置容器启动时执行的命令
#CMD ["./server"]
#ENTRYPOINT ./server -c config.docker.yaml

CMD ["sh", "-c", "./server -c config.docker.yaml"]
