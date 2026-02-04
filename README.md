# Wire Learn 教程

一个最小的 Go + Gin + GORM 教程项目，演示如何使用 Google Wire 做依赖注入。包含基于 SQLite 的简单 User API。

## 技术栈

- Go
- Gin
- GORM（SQLite）
- Google Wire
- Zap（日志）

## 目录结构

- api/: HTTP 接口层
- core/: 通用 provider（DB、Logger）
- model/: 数据模型
- router/: 路由注册
- service/: 业务逻辑
- wire/: Wire 注入器

## 前置条件

- Go 1.20+（或本地版本）
- 安装 Wire: `go install github.com/google/wire/cmd/wire@latest`

## 安装依赖

```bash
go mod tidy
```

## 分步教程

该项目展示了从全局变量到手动 wiring，再到 Google Wire 的演进过程。

### 1) 全局变量（baseline）

目标：快速跑通。

- 在 `core/` 中创建全局单例（DB、Logger）。
- 在 service / api 层直接使用。

示意（仅概念）：

- `core` 暴露 `DB` 与 `Logger`
- `service` 和 `api` 直接使用 `core.DB` / `core.Logger`

缺点：耦合高、难测试、依赖不显式。

### 2) 手动 wiring 去除全局变量

目标：让依赖显式、可测试。

- 增加构造函数：
  - `core.NewDB()`
  - `core.NewLogger()`
  - `service.NewUserService(db, logger)`
  - `api.NewUserApi(userService)`
  - `router.NewUserRouter(api)`
- 在 `main.go` 中手动拼装依赖。

这样每一层依赖都来自参数（无全局变量）。

### 3) 引入 Google Wire 自动注入

目标：避免重复手动 wiring。

1. 在 `wire/wire.go` 创建 Wire 注入器（build tag: `wireinject`）。
2. 在 `wire.Build(...)` 中注册 provider：
   - `core.NewDB`
   - `core.NewLogger`
   - `service.NewUserService`
   - `api.NewUserApi`
   - `router.NewUserRouterProvider`
3. 运行 Wire 生成 `wire_gen.go`：

```bash
wire ./wire
```

Wire 会生成完整依赖图，并提供 `InitWire()` 返回 `*gin.Engine`。

### 4) 使用生成的注入器

- 在 `main.go` 中调用 `wire.InitWire()` 替代手动 wiring。

至此，依赖显式、可测试，并自动装配。

## 日志（Logger）

- 日志由 Zap 提供，在 `core/` 中创建并注入到 `service` 与 `api`。
- 推荐在 API 层记录请求入口，在 Service 层记录业务与错误。
- 典型写法：
  - `logger.Info("message")`
  - `logger.Error("message", zap.Error(err))`

## 生成 Wire

```bash
wire ./wire
```

## 运行

```bash
go run .
```

## API

### 获取用户

`GET /api/user?id=1`

示例：

```bash
curl "http://localhost:8080/api/user?id=1"
```

## 备注

- SQLite 数据库文件：`test.db`
- 新增 service 时记得在 `wire/` 中补充 provider。
