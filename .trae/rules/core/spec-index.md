---
trigger: manual
---

# 规则索引（go-layout）v1.6

默认：
- Profile：GoService

仓库识别（以文件存在性判断）：
- 微服务仓库：存在 `buf.gen.yaml`（生成配置）
- Proto 仓库：存在 `buf.yaml` 或 `buf.work.yaml`（模块/工作区配置）
- 若同时存在：视为 monorepo；以实际 `buf` 配置与生成产物目录为准

Proto 协作上下文：
- 微服务仓库通常不包含 `.proto`，只维护生成产物（常见为 `dep/protobuf/gen`）
- Proto 定义通常维护在独立的 Proto 仓库中，通过 `buf generate` 同步到微服务仓库

快速用法：
- 第一步：确认仓库类型（见“仓库识别”）
- 第二步：按变更类型跳转规则
  - 新增/修改 Proto：`architecture/proto-spec.md` + `architecture/response-spec.md` + `architecture/status-code-spec.md` + `architecture/validation-spec.md`
  - 新增/修改 Service RPC：`architecture/service-spec.md` + `quality/error-handling-spec.md` + `quality/testing-spec.md`
  - 新增/修改 Biz：`core/requirements-spec.md` + `quality/testing-spec.md` + `architecture/comment-spec.md`
  - 新增/修改 Data：`architecture/data-spec.md` + `quality/testing-spec.md` + `architecture/comment-spec.md`
  - 仅调整生成链路：`core/workflow-spec.md` + `architecture/buf-spec.md`
- 第三步：按“执行清单（按变更触发）”做最小验证

Profile：
- GoService：core 三件套（requirements/workflow/naming）
- GoServiceStrict：在 GoService 基础上额外加载 architecture 与 quality（API/错误/安全/测试）

规则约束：
- 规则文件单篇不超过 1000 字

交付校验：
- 按 `quality/testing-spec.md` 的“按场景验收”规则执行

执行清单（按变更触发）：
- 仅改规则/文档：检查字数限制与引用一致性即可
- 变更 Go 代码：必须通过 `go test ./...`（建议同时 `go vet ./...`）
- 变更 Proto 定义：在 Proto 仓库执行 `buf lint` → `buf push`，并在微服务仓库执行 `buf generate`（或 `make generate/make init`）
- 变更生成相关配置（如 `buf.gen.yaml`/wire/goverter 配置）：必须重新生成并确保编译/测试通过

对外输出底线：
- 真实密钥/令牌/密码/个人敏感信息不得进入代码与仓库（见 `quality/security-spec.md`）

加载顺序：
- GoService：
  - core/requirements-spec.md
  - core/workflow-spec.md
  - core/naming-conventions.md
  - core/doc-url-spec.md
- GoServiceStrict：在 GoService 基础上额外加载
  - architecture/api-design-spec.md（API 规范索引）
  - architecture/*-spec.md（API/Proto/Response/Code/Service/Data/Style/Comment/Context/Validation）
  - architecture/protovalidate-cheatsheet.md（校验速查）
  - quality/*.md（error-handling/security/testing）
