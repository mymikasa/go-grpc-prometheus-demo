FROM prom/prometheus:latest

ENV PROMETEUS_TARGET prometheus:8080
ENV GRPC_SERVER_TARGET grpcserver:8080
ENV GRPC_CLIENT_TARGET grpcclient:8081

RUN sed -i -e "s/{ENV_PROMETHEUS_TARGET}/$PROMETEUS_TARGET/" -e "s/{ENV_GRPC_SERVER_TARGET}/$GRPC_SERVER_TARGET/" -e "s/{ENV_GRPC_CLIENT_TARGET}/$GRPC_CLIENT_TARGET/"

# 复制 Prometheus 配置文件到容器中
COPY prometheus.yml /etc/prometheus/prometheus.yml

# 暴露 Prometheus 的 Web 接口端口
EXPOSE 9090

# 指定启动命令，指定配置文件路径
ENTRYPOINT [ "/bin/prometheus" ]
CMD        [ "--config.file=/etc/prometheus/prometheus.yml" ]
