# DAuth 项目架构文档

## 项目概述

DAuth 是一个基于 Go 语言开发的认证服务，采用 Clean Architecture（整洁架构）模式设计，提供用户认证、身份管理、OAuth2、会话管理等核心功能。

### 技术栈

- **语言**: Go 1.25.5
- **框架**: Ellie（微服务框架）
- **传输层**: Ellie HTTP、Ellie gRPC
- **数据库**: MySQL（GORM）
- **缓存**: Redis
- **认证**: JWT
- **服务注册与发现**: Consul
- **链路追踪**: OpenTelemetry
- **依赖注入**: Wire
- **代码生成**: Buf（Protocol Buffer）

## 架构设计

### Clean Architecture（整洁架构）

项目采用整洁架构，代码分层清晰，依赖方向从外向内：

```
┌─────────────────────────────────────────────────────────────┐
│                        Handler Layer                         │
│  (HTTP/gRPC 处理层，负责请求响应转换和调用 Application)      │
└────────────────────────────┬────────────────────────────────┘
                             │
┌────────────────────────────▼────────────────────────────────┐
│                      Application Layer                       │
│  (应用层，编排用例，调用 Domain 层的 Biz)                   │
└────────────────────────────┬────────────────────────────────┘
                             │
┌────────────────────────────▼────────────────────────────────┐
│                       Domain Layer                           │
│  (领域层，包含业务逻辑 Entity、Biz、DTO、Repo 接口)         │
└────────────────────────────┬────────────────────────────────┘
                             │
┌────────────────────────────▼────────────────────────────────┐
│                     Infrastructure Layer                     │
│  (基础设施层，实现 Repo 接口，数据库、缓存、外部服务调用)     │
└─────────────────────────────────────────────────────────────┘
```

## 核心模块

### 1. API 层 (`api/`)

- **proto/**: Protocol Buffer 定义文件，定义 gRPC 服务和消息
  - `authn/`: 认证相关服务定义
  - `identity/`: 身份相关服务定义
  - `oauth2/`: OAuth2 相关服务定义
  - `session/`: 会话相关服务定义
  - `common/`: 通用类型定义
  - `errdef/`: 错误定义
- **gen/**: 生成的代码
  - 由 Buf 工具从 proto 生成的 Go 代码
  - HTTP 网关代码
  - 错误定义代码

### 2. Command 层 (`cmd/`)

- **dauth-cli/**: 命令行工具
- **wire.go/wire_gen.go**: Wire 依赖注入配置
- **run.go**: 服务启动入口

### 3. 内部模块 (`internal/`)

#### 3.1 Domain 层 (`internal/domain/`)

领域层是业务逻辑的核心，定义了实体、业务逻辑和仓库接口：

##### 认证模块 (`authn/`)
- **biz/authn.go**: 认证业务逻辑（登录、登出、认证状态检查）
- **dto/authn.go**: 认证数据传输对象
- **repo/repo.go**: 认证仓库接口

##### 身份模块 (`identity/`)
- **biz/user.go**: 用户业务逻辑（创建用户、查询用户、验证密码）
- **entity/user.go**: 用户实体
- **dto/user.go**: 用户数据传输对象
- **repo/repo.go**: 用户仓库接口

##### OAuth2 模块 (`oauth2/`)
- **biz/oauth.go**: OAuth2 业务逻辑
- **repo/repo.go**: OAuth2 仓库接口

##### 会话模块 (`session/`)
- **biz/session.go**: 会话业务逻辑
- **repo/repo.go**: 会话仓库接口

#### 3.2 Application 层 (`internal/application/`)

应用层负责用例编排，调用 Domain 层的 Biz：

- **application.go**: Application Provider Set 定义
- **authn.go**: 认证应用服务
- **identity.go**: 身份应用服务
- **oauth2.go**: OAuth2 应用服务
- **session.go**: 会话应用服务

#### 3.3 Handler 层 (`internal/handler/`)

处理层负责 HTTP/gRPC 请求处理和响应转换：

- **handler.go**: Handler Provider Set 定义
- **authn.go**: 认证处理
- **identity.go**: 身份处理
- **oauth2.go**: OAuth2 处理
- **session.go**: 会话处理

#### 3.4 Infrastructure 层 (`internal/infra/`)

基础设施层负责技术实现，包括数据库、缓存、外部服务等：

##### Foundation (`foundation/`)
- **db.go**: 数据库初始化
- **redis.go**: Redis 初始化
- **logger.go**: 日志初始化
- **tracer.go**: 链路追踪初始化
- **registry.go**: 服务注册初始化

##### Persistence (`persistence/`)
- **core/**: 核心数据访问层，使用 GORM
- **impl/authn/**: 认证仓库实现
- **impl/identity/**: 身份仓库实现
- **impl/oauth2/**: OAuth2 仓库实现
- **impl/session/**: 会话仓库实现
- **migration/**: 数据库迁移脚本
  - `identity.sql`: 身份相关表结构
  - `authn.sql`: 认证相关表结构

##### Cache (`cache/`)
- **cache.go**: 缓存接口
- **impl/jwt_cache.go**: JWT 缓存实现

##### RPC (`rpc/`)
- **dauth/authn.go**: 认证 RPC 调用
- **dauth/identity.go**: 身份 RPC 调用

##### Utils (`utils/`)
- **security/jwt/**: JWT 工具
- **security/encryption.go**: 加密工具
- **ctxutil/**: 上下文工具
- **cache/**: 缓存工具

#### 3.5 Server (`internal/server/`)

- **http.go**: Ellie HTTP 服务器配置
- **grpc.go**: Ellie gRPC 服务器配置
- **middleware/**: 中间件
  - `jwt_auth.go`: JWT 认证中间件
  - `auth_white_list.go`: 白名单配置

**服务注册**: 通过 `ServiceRegistrar` 将 Handler 注册到 Ellie 服务器，同时支持 HTTP 和 gRPC 两种协议

#### 3.6 Config (`internal/conf/`)

- **app_config.go**: 应用配置定义
- **conf.go**: 配置加载

### 4. 配置文件 (`configs/`)

- **config.toml**: 基础配置
- **config.dev.toml**: 开发环境配置
- **config.ppe.toml**: 预发布环境配置
- **config.prod.toml**: 生产环境配置

### 5. 脚本 (`scripts/`)

- **Dockerfile**: Docker 镜像构建文件
- **build_docker.sh**: Docker 构建脚本
- **generate_tls_cert.sh**: TLS 证书生成脚本

## 工作流程示例

### 用户登录流程

1. **Handler 层**: HTTP 请求到达 `handler/authn.go` 的 Login 方法
2. **Application 层**: 调用 `application/authn.go` 的 Login 方法
3. **Domain 层**: 调用 `domain/authn/biz/authn.go` 的 Login 方法
   - 通过 RPC 调用身份模块验证密码
   - 生成 JWT Token
4. **Response**: 返回 Token 给客户端

### 创建用户流程

1. **Handler 层**: HTTP 请求到达 `handler/identity.go` 的 CreateUser 方法
2. **Application 层**: 调用 `application/identity.go` 的 CreateUser 方法
3. **Domain 层**: 调用 `domain/identity/biz/user.go` 的 CreateUser 方法
   - 加密密码
   - 创建用户实体
4. **Infrastructure 层**: 调用 `infra/persistence/impl/identity/users.go` 保存到数据库

## 核心依赖注入

项目使用 Wire 进行依赖注入，主要在以下文件定义：

- `cmd/wire.go`: Wire 配置
- `cmd/wire_gen.go`: 生成的依赖注入代码
- `internal/domain/domain.go`: Domain 层 Provider Set
- `internal/application/application.go`: Application 层 Provider Set

## 数据库表结构

### Identity 模块表

- **identity_users**: 用户表
  - uid: 用户ID
  - username: 用户名
  - password: 密码（加密）
  - nickname: 昵称
  - phone: 手机号
  - email: 邮箱
  - avatar: 头像
  - status: 状态
  - created_at/updated_at: 时间戳

### Authn 模块表

- **authn_policies**: 认证策略表
- **authn_policy_targets**: 认证策略目标表
- **authn_attempts**: 认证尝试记录表

## 开发指南

### 代码生成

```bash
# 生成 Protocol Buffer 代码
buf generate

# 生成 Wire 依赖注入代码
wire ./cmd

# 生成 GORM DAO 代码
# 查看 cmd/dauth-cli/cmd/gen.go
```

### 运行服务

```bash
go run main.go run
```

### 配置

配置文件位于 `configs/` 目录，支持多环境配置。

### 代码规范

- 使用 Go modules 管理依赖
- 遵循 Go 风格指南（`go fmt`）
- 使用 Wire 进行依赖注入
- 使用 Protocol Buffers 定义 API
- 使用 gRPC gateway 提供 HTTP 支持

## 关键文件速查

| 功能              | 文件路径                                 |
| ----------------- | ---------------------------------------- |
| 用户实体          | internal/domain/identity/entity/user.go  |
| 用户业务逻辑      | internal/domain/identity/biz/user.go     |
| 认证业务逻辑      | internal/domain/authn/biz/authn.go       |
| JWT 工具          | internal/infra/utils/security/jwt/jwt.go |
| 数据库迁移        | internal/infra/persistence/migration/    |
| 配置加载          | internal/conf/conf.go                    |
| Ellie HTTP 服务器 | internal/server/http.go                  |
| Ellie gRPC 服务器 | internal/server/grpc.go                  |
| 服务注册          | internal/handler/handler.go              |
