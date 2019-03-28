# GP Server Part

+ 所有模块都定义成接口，方便替换
+ web节点、判题节点分离
+ 支持docker部署

## 模块定义（./module）

### 底层模块
+ 网关（gate.go）
+ 项目配置 (config.go)
+ 数据库（db.go）
+ 共享缓存（cache.go）

### 业务逻辑模块
+ 用户认证 (auth.go)
+ 管理员管理 (manage.go)



