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
