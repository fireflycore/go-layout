---
trigger: manual
---

# 规则索引（go-layout）v1.2

默认：
- Profile：GoService
- Proto 仓库：优先按仓库现状自动发现（`buf.yaml/buf.work.yaml`）
  - 若本仓库不包含 `.proto`：默认视为“Proto 独立仓库 + 本仓库生成产物（dep/protobuf/gen）”
  - 若无法自动发现：按本仓库同级目录的 `firefly/` 作为常见默认值（仅描述约定，不强依赖绝对路径）

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
