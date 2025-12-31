# Trae 规则（go-layout）

本目录用于承载本项目的通用 AI 开发规则集，目标是让 AI 的输出与 `go-layout` 的工程化实践保持一致：分层架构、Proto First、Wire 依赖注入、配置加载方式、日志与错误处理风格，以及本项目的目录结构约束。

这些规则与项目文档配套使用：
- `docs/project-guide.md`：工具链与快速开始
- `docs/directory-structure.md`：目录与分层边界
- `docs/architecture.md`：Wire 与依赖方向
- `docs/data-flow.md`：请求链路与读写策略
- `docs/best-practices.md`：新增模块 checklist 与约定

## 推荐使用方式

- 以 `core/spec-index.zh-CN.md` 作为入口，按需补充引用其它模块
- 在实现需求前先遵循：分层边界（Server/Service/Biz/Data）、Proto First、Wire 注册与生成、Goverter/Buf 生成产物位置

## 📦 目录结构

```
.trae/
├── rules/
│   ├── core/                           # 核心规范（必需）
│   │   ├── spec-index.md
│   │   ├── spec-index.zh-CN.md
│   │   ├── requirements-spec.md
│   │   ├── requirements-spec.zh-CN.md
│   │   ├── workflow-spec.md
│   │   ├── workflow-spec.zh-CN.md
│   │   ├── naming-conventions.md
│   │   └── naming-conventions.zh-CN.md
│   ├── architecture/                  # 架构与接口规范（按需）
│   │   └── api-design-spec.zh-CN.md
│   └── quality/                       # 质量保证规范（推荐）
│       ├── security-spec.zh-CN.md
│       ├── error-handling-spec.zh-CN.md
│       └── testing-spec.zh-CN.md
└── README.md
```
