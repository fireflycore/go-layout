---
trigger: manual
---

# 规则索引（go-layout）v1.1

默认：
- Profile：GoService
- Proto 仓库：优先自动发现 `buf.yaml/buf.work.yaml`；fallback：`/Users/lhdht/product/lhdht/code/firefly`

Profile：
- GoService：core 三件套（requirements/workflow/naming）
- GoServiceStrict：在 GoService 基础上额外加载 architecture 与 quality

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
  - architecture/api-design-spec.md（索引）
  - architecture/*-spec.md
  - quality/*.md
