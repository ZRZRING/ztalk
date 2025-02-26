# 在线论坛系统后端

下载 redis 和 mysql 并启动，修改 config.yaml 的相关配置

启动服务

```bash
go mod tidy
go run main.go config/dev.yaml
```

使用 [air](https://github.com/air-verse/air) 可以进行热部署

```bash
go install github.com/air-verse/air@latest
air
```