# Rope Net 🕸️

> 基于 Go 语言开发的社区论坛与任务管理综合平台

一个集成了社区论坛、任务管理和实时通信功能的现代化 Web 应用，支持用户发布帖子、实时评论、任务提醒等功能。

## ✨ 功能特性

### 🔐 用户系统
- **双步验证登录**：密码验证 + 邮箱验证码确保账户安全
- **Token 认证**：24 小时有效期的自动过期机制
- **会话管理**：基于 Gin Session 的会话保持

### 📝 论坛系统
- **帖子管理**：发布、查看、删除帖子
- **浏览量统计**：自动记录帖子浏览次数
- **热度排序**：按浏览量倒序展示帖子列表
- **实时评论**：基于 WebSocket 的实时评论同步
- **线程评论**：支持评论的多级回复

### ✅ 任务系统
- **任务管理**：创建、完成、删除个人任务
- **截止日期**：为任务设置 deadline
- **实时提醒**：后台定时检查，WebSocket 推送过期提醒
- **任务状态**：已完成/未完成状态管理

### 🔌 实时通信
- **评论 WebSocket**：实时同步帖子评论和删除操作
- **任务 WebSocket**：实时推送任务过期提醒
- **历史数据回放**：新连接自动接收历史评论

## 🛠️ 技术栈

### 后端框架
- **Go** 1.23.5 - 高性能编程语言
- **Gin** v1.10.0 - 轻量级 Web 框架
- **GORM** v1.25.12 - ORM 框架
- **Gorilla WebSocket** v1.5.3 - WebSocket 支持

### 数据存储
- **MySQL** - 关系型数据库
- **go-cache** v2.1.0 - 内存缓存（验证码存储）

### 其他组件
- **QQ SMTP** - TLS 加密邮件服务
- **Gin Sessions** - 会话管理
- **标准库 Logger** - 日志记录

### 预留组件
- **RabbitMQ** v1.1.0 - 消息队列（预留）
- **Kafka** v0.4.47 - 消息流（预留）

## 🚀 快速开始

### 环境要求

- Go 1.23.5 或更高版本
- MySQL 5.7 或更高版本
- QQ 邮箱（用于发送验证码）

### 安装步骤

1. **克隆项目**
```bash
git clone <repository-url>
cd "Rope Net"
```

2. **安装依赖**
```bash
go mod download
```

3. **配置数据库**

编辑 `config/db_config.json`：
```json
{
  "database": {
    "type": "mysql",
    "host": "127.0.0.1",
    "port": 3306,
    "user": "your_username",
    "password": "your_password",
    "dbname": "rope_net"
  }
}
```

创建数据库：
```sql
CREATE DATABASE rope_net CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

4. **配置邮箱**

编辑 `config/qq_email_config.json`：
```json
{
  "email": {
    "host": "smtp.qq.com",
    "port": 465,
    "username": "your_qq_email@qq.com",
    "password": "your_qq_authorization_code"
  }
}
```

> 💡 提示：QQ 邮箱授权码获取方式：QQ 邮箱 → 设置 → 账户 → POP3/IMAP/SMTP/Exchange/CardDAV/CalDAV 服务 → 生成授权码

5. **数据库迁移**

程序首次运行会自动创建表结构（GORM AutoMigrate）

6. **运行程序**
```bash
go run cmd.go
```

服务将在 `http://localhost:8080` 启动

## 📁 项目结构

```
Rope Net/
├── api/                          # API 接口层
│   ├── handlers/                 # 请求处理程序
│   │   ├── comment_handlers/     # 评论相关处理
│   │   ├── post_handlers/        # 帖子相关处理
│   │   ├── task_handlers/        # 任务相关处理
│   │   └── user_handlers/        # 用户相关处理
│   └── routes/                   # 路由定义
│       └── routes.go             # 所有 API 路由
├── config/                       # 配置文件
│   ├── db_config.json           # 数据库配置
│   └── qq_email_config.json     # 邮箱配置
├── internal/                     # 内部配置管理
│   ├── dbConfig.go              # 数据库配置读取
│   └── qqEmailConfig.go         # 邮箱配置读取
├── middleware/                   # 中间件
│   └── identifyTokenMiddleware.go # Token 验证中间件
├── models/                       # 数据模型
│   ├── users.go                 # 用户模型
│   ├── posts.go                 # 帖子模型
│   ├── tasks.go                 # 任务模型
│   ├── post_comments.go         # 帖子评论模型
│   └── thread_comments.go       # 线程评论模型
├── pkg/                          # 工具包
│   ├── db/                       # 数据库连接
│   ├── logger/                   # 日志工具
│   ├── identify/                 # 身份认证
│   │   ├── token/               # Token 生成与验证
│   │   └── verification_code/   # 验证码生成与发送
│   └── rabbitmq/                # RabbitMQ 连接（预留）
├── cmd.go                        # 程序入口
├── go.mod                        # Go 模块定义
├── go.sum                        # 依赖版本锁定
└── README.md                     # 项目说明
```

## 📡 API 文档

### 用户接口

#### 注册
```http
POST /api/user/register
Content-Type: application/json

{
  "username": "string",
  "email": "string",
  "password": "string"
}
```

#### 预登录（发送验证码）
```http
POST /api/user/preLogin
Content-Type: application/json

{
  "username": "string",
  "password": "string"
}
```

#### 最终登录（验证验证码）
```http
POST /api/user/finalLogin
Content-Type: application/json

{
  "username": "string",
  "verificationCode": "string"
}

Response:
{
  "token": "string"
}
```

### 帖子接口

#### 发布帖子 🔒
```http
POST /api/post/publish
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "string",
  "content": "string"
}
```

#### 获取帖子
```http
GET /api/post/getPost/:postID
```

#### 获取帖子列表
```http
GET /api/post/getPostList
```

#### 删除帖子 🔒
```http
DELETE /api/post/deletePost/:postID
Authorization: Bearer <token>
```

### 评论接口

#### WebSocket 连接 🔒

**连接地址**：
```
ws://localhost:8080/api/comment/ws?postID=<postID>
```

**请求头**：
```json
{
  "Authorization": "Bearer <token>"
}
```

**发送评论消息**：
```json
{
  "userID": 1,
  "postID": 1,
  "content": "评论内容"
}
```

**接收广播消息**：
```json
{
  "id": 1,
  "userID": 1,
  "postID": 1,
  "content": "评论内容",
  "createTime": "2025-12-12T10:00:00Z"
}
```

#### 创建线程评论 🔒
```http
POST /api/comment/createThreadComment
Authorization: Bearer <token>
Content-Type: application/json

{
  "commentID": 1,
  "content": "string"
}
```

#### 获取线程评论
```http
GET /api/comment/getThreadComment/:commentID
```

#### 删除帖子评论 🔒
```http
DELETE /api/comment/deletePostComment/:postCommentID
Authorization: Bearer <token>
```

#### 删除线程评论 🔒
```http
DELETE /api/comment/deleteThreadComment/:threadCommentID
Authorization: Bearer <token>
```

### 任务接口

#### 创建任务 🔒
```http
POST /api/task/createTask
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "string",
  "deadline": "2025-12-31T23:59:59Z"
}
```

#### 获取任务列表 🔒
```http
GET /api/task/getTask
Authorization: Bearer <token>
```

#### 完成任务 🔒
```http
POST /api/task/completeTask/:taskID
Authorization: Bearer <token>
```

#### 删除任务 🔒
```http
DELETE /api/task/deleteTask/:taskID
Authorization: Bearer <token>
```

#### WebSocket 连接（任务提醒）🔒

**连接地址**：
```
ws://localhost:8080/api/task/ws
```

**请求头**：
```json
{
  "Authorization": "Bearer <token>"
}
```

**接收提醒消息**：
```json
{
  "type": "reminder",
  "taskID": 1,
  "taskName": "任务名称",
  "deadline": "2025-12-31T23:59:59Z",
  "message": "任务已过期"
}
```

> 🔒 标记的接口需要在请求头中携带 `Authorization: Bearer <token>`

## ⚙️ 配置说明

### 数据库配置

`config/db_config.json`：
```json
{
  "database": {
    "type": "mysql",
    "host": "127.0.0.1",
    "port": 3306,
    "user": "root",
    "password": "password",
    "dbname": "rope_net"
  }
}
```

**字段说明**：
- `type`：数据库类型
- `host`：数据库地址
- `port`：数据库端口
- `user`：数据库用户名
- `password`：数据库密码
- `dbname`：数据库名称

### 邮箱配置

`config/qq_email_config.json`：
```json
{
  "email": {
    "host": "smtp.qq.com",
    "port": 465,
    "username": "your_email@qq.com",
    "password": "authorization_code"
  }
}
```

**字段说明**：
- `host`：SMTP 服务器地址
- `port`：SMTP 端口（TLS 加密）
- `username`：QQ 邮箱地址
- `password`：QQ 邮箱授权码

## 🔧 核心功能实现

### 身份认证流程

```
用户注册 → 数据库存储
↓
预登录(username, password) → 验证密码 → 生成验证码 → 发送邮件 → 缓存验证码(5分钟)
↓
最终登录(verificationCode) → 验证缓存 → 生成Token(16字符随机) →
→ 保存到数据库(24小时过期) → 返回授权头
↓
后续请求 → Authorization: Bearer <token> → 中间件验证 → token有效期检查
```

### WebSocket 实时通信

**评论系统**：
- 用户连接时自动发送历史评论
- 新评论通过 WebSocket 实时广播给同一帖子的所有在线用户
- 评论删除时同步通知所有连接的客户端

**任务提醒**：
- 后台每分钟检查一次所有未完成任务
- 发现过期任务后通过 WebSocket 推送给对应用户
- 用户可实时收到任务过期提醒

### 验证码机制

- 验证码有效期：5 分钟
- 存储方式：内存缓存（go-cache）
- 发送方式：QQ SMTP TLS 加密邮件
- 验证码格式：6 位随机数字

## 🗄️ 数据库表结构

### users 表
```sql
CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    token VARCHAR(255),
    token_expires_at DATETIME
);
```

### posts 表
```sql
CREATE TABLE posts (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    page_view INT DEFAULT 0,
    publish_time DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

### post_comments 表
```sql
CREATE TABLE post_comments (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    post_id BIGINT UNSIGNED NOT NULL,
    content TEXT NOT NULL,
    create_time DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (post_id) REFERENCES posts(id)
);
```

### thread_comments 表
```sql
CREATE TABLE thread_comments (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    comment_id BIGINT UNSIGNED NOT NULL,
    content TEXT NOT NULL,
    create_time DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (comment_id) REFERENCES post_comments(id)
);
```

### tasks 表
```sql
CREATE TABLE tasks (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    name VARCHAR(255) NOT NULL,
    deadline DATETIME NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

## 🎯 路线图

### 已完成 ✅
- [x] 用户注册与登录
- [x] 双步验证（密码 + 邮箱验证码）
- [x] Token 认证机制
- [x] 帖子发布与管理
- [x] 实时评论系统
- [x] 任务管理与提醒
- [x] WebSocket 实时通信
- [x] 浏览量统计

### 计划中 📋
- [ ] 密码加密存储（bcrypt）
- [ ] 环境变量管理（避免配置文件硬编码）
- [ ] 请求参数验证
- [ ] 更详细的错误处理
- [ ] 用户头像上传
- [ ] 帖子图片上传
- [ ] 点赞/收藏功能
- [ ] 用户关注系统
- [ ] 消息通知系统
- [ ] 搜索功能
- [ ] 分页优化
- [ ] RabbitMQ 集成
- [ ] Kafka 集成
- [ ] 单元测试
- [ ] API 文档（Swagger）
- [ ] Docker 支持

## 🔒 安全性建议

> ⚠️ **重要提示**：当前版本存在以下安全问题，建议在生产环境部署前改进：

1. **密码存储**：密码应使用 bcrypt 等哈希算法加密，而非明文存储
2. **配置安全**：数据库密码和邮箱授权码应使用环境变量，避免在配置文件中硬编码
3. **输入验证**：添加更严格的输入参数验证，防止 SQL 注入等攻击
4. **HTTPS**：生产环境应使用 HTTPS 加密传输
5. **跨域配置**：合理配置 CORS 策略
6. **限流保护**：添加 API 限流机制，防止恶意请求

## 🤝 贡献指南

欢迎贡献代码！请遵循以下步骤：

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 提交 Pull Request

### 代码规范

- 遵循 Go 官方代码规范
- 使用 `gofmt` 格式化代码
- 添加必要的注释
- 编写单元测试

## 📄 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 📧 联系方式

如有问题或建议，欢迎通过以下方式联系：

- 提交 Issue
- 发送邮件至：[your-email@example.com]

## 🙏 致谢

感谢以下开源项目：

- [Gin](https://github.com/gin-gonic/gin) - Web 框架
- [GORM](https://gorm.io/) - ORM 框架
- [Gorilla WebSocket](https://github.com/gorilla/websocket) - WebSocket 库
- [go-cache](https://github.com/patrickmn/go-cache) - 内存缓存

---

⭐ 如果这个项目对你有帮助，请给我们一个 Star！
