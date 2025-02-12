# Rope Net

Rope Net 是一个基于 Go 语言开发的后端项目，能实现发布帖子，可以进行实时评论，有浏览量设置，以此可以推送浏览量高的帖子，另有任务板块，可以记录个人一段时间内的任务，可以查看完成情况，可以设置期限，有时间截止提醒等。

## 项目结构

- `cmd.go` 是项目的入口文件。
- `go.mod` 和 `go.sum` 是 Go 模块的配置文件。
- `api` 目录包含与 API 相关的代码，如路由和处理程序。
- `pkg` 目录包含一些通用的包。
- `config` 目录包含项目的配置文件。
- `middleware` 目录包含中间件代码。
- `models` 目录包含数据模型的定义。
- `internal` 目录包含项目内部使用的代码。

## 主要依赖

- Gin：用于构建 Web 服务器和处理路由。
- GORM：用于数据库操作。
- WebSocket：用于要求实时相关操作。

## 快速开始

1. 克隆项目：

   

   ```sh
   git clone https://github.com/xkiven/rope_net/tree/master
   cd Rope_Net
   ```

2. 配置文件：
   确保 `config/qq_email_config.json` 文件中 QQ 邮箱配置信息正确。

3. 配置数据库：
   项目使用的数据库配置在 `pkg/db/connectDB.go` 中读取，你需要确保相关数据库配置正确。

4. 运行项目：

   收起

   

   ```sh
   go run cmd.go  # 假设你的入口文件是 cmd.go
   ```

5. 访问 API：
   项目默认使用 Gin 框架运行在某个端口（通常为 `:8080`），你可以使用工具如 Apifox 来测试 API。

## 注意事项

请确保你的系统已经安装了 Go 语言环境，并且数据库服务正常运行。
