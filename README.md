# Firefly Go Layout

`go-layout` 是 Firefly 微服务框架的 Go 版本标准项目模板。它提供了一套标准化的目录结构和基础设施配置，旨在帮助开发者快速构建规范的微服务应用。

本模板基于 **[go-micro](github.com/fireflycore/go-micro)**（Firefly 微服务框架的 Go 版本核心库）构建。

> [在线文档](https://firefly.lhdht.cn/guide/)

## 快速开始

### 初始化项目

1.  **Clone 项目**

    ```bash
    git clone https://github.com/fireflycore/go-layout.git my-project
    cd my-project
    ```

2.  **重命名模块**

    使用提供的脚本将模块名（默认 `go-layout`）替换为你自己的模块名（例如 `github.com/myuser/my-project`）。

    **Windows:**
    ```cmd
    .\rename_project.bat github.com/myuser/my-project
    ```

    **Linux / macOS:**
    ```bash
    chmod +x rename_project.sh
    ./rename_project.sh github.com/myuser/my-project
    ```

3.  **整理依赖**

    ```bash
    go mod tidy
    ```

### 常用命令

- `make run`: 运行服务
- `make build`: 编译服务
