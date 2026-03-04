# 使用 CentOS 7.9 作为基础镜像
FROM centos:7

# 接收 Jenkins 传入的变量（build-arg）
ARG DOCKER_IMAGE

# 创建一个目录用于存放 Go 二进制文件
WORKDIR /app

# 从 deploy 目录拷贝与 DOCKER_IMAGE 同名的二进制
# （对应 Jenkins 的 go build -o ./deploy/<svc>）

COPY deploy/${DOCKER_IMAGE} /app/

# 赋予文件执行权限
RUN chmod +x /app/${DOCKER_IMAGE} \
    # 建一个固定名的软链，避免 CMD 里变量展开的坑
    && ln -sf /app/${DOCKER_IMAGE} /app/app

# 配置文件路径（如需可再做成 ARG -> ENV）
ENV CONFIG_PATH=/opt/app.yaml

# 用固定名启动，避免 JSON 形式的 CMD 不展开变量
CMD ["/app/app"]
