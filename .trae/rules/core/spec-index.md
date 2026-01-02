---
trigger: manual
---

# 规则索引（go-layout）v1.2

默认：
- Profile：GoService

仓库识别（以文件存在性判断）：
- 微服务仓库：存在 `buf.gen.yaml`（生成配置）
- Proto 仓库：存在 `buf.yaml` 或 `buf.work.yaml`（模块/工作区配置）
- 若同时存在：视为 monorepo；以实际 `buf` 配置与生成产物目录为准

Proto 协作上下文：
- 微服务仓库通常不包含 `.proto`，只维护生成产物（常见为 `dep/protobuf/gen`）
- Proto 定义通常维护在独立的 Proto 仓库中，通过 `buf generate` 同步到微服务仓库

Profile：
- GoService：core 三件套（requirements/workflow/naming）
- GoServiceStrict：在 GoService 基础上额外加载 architecture 与 quality（API/错误/安全/测试）

规则约束：
- 规则文件单篇不超过 1000 字
- 存量兼容：非需求驱动不做“顺手对齐”（字段序/命名/结构），新写/新改按最新规则对齐

交付校验：
- 按 `quality/testing-spec.md` 的“按场景验收”规则执行

执行清单（按变更触发）：
- 仅改规则/文档：检查字数限制与引用一致性即可
- 变更 Go 代码：必须通过 `go test ./...`（建议同时 `go vet ./...`）
- 变更 Proto 定义：在 Proto 仓库执行 `buf lint`，并在微服务仓库执行 `buf generate`（或 `make generate/make init`）
- 变更生成相关配置（如 `buf.gen.yaml`/wire/goverter 配置）：必须重新生成并确保编译/测试通过

对外输出底线：
- 响应体/日志/错误信息不得包含敏感信息（见 `quality/security-spec.md`）
- 存量兼容优先级高于“顺手优化”：仅在需求触发的文件/接口上对齐规则

加载顺序：
- GoService：
  - core/requirements-spec.md
  - core/workflow-spec.md
  - core/naming-conventions.md
- GoServiceStrict：在 GoService 基础上额外加载
  - architecture/api-design-spec.md（API 规范索引）
  - architecture/*-spec.md（API/Proto/Response/Code/Service/Data/Style/Comment/Context/Validation）
  - architecture/protovalidate-cheatsheet.md（校验速查）
  - quality/*.md（error-handling/security/testing）
