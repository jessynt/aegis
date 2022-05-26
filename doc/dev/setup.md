# 项目开发设定

```
git clone ${GIT_REPO_PATH} /path/to/workspace
```

## 依赖管理

目前使用 go module 对依赖进行管理

## 测试

```
make test       # 执行单元测试
make test-bench # 执行压力测试
make test-race  # 执行竞态测试
make test-lint  # 执行风格测试
```

## 构建

```
make binary
```

## 其他

其他开发命令可以参考 `Makefile` 内容。
