# API 服务
go build -o ./deploy/api ./cmd/api/main.go
## 启动
docker compose -f /home/ubuntu/sf/docker-compose.yml  -p "sf-api" up -d api
## 查看日志
docker logs --tail=1000 -f sf-api
