# Trae 规则（go-layout）

本目录用于约束 AI 的代码输出与当前微服务项目保持一致：分层边界、Proto First、生成链路（buf/goverter/wire）、日志与错误处理风格。
代码与 Proto 需要补充必要注释：优先解释业务语义、约束与边界，避免逐行翻译。

推荐在对话中仅引用入口文件（需要更严格时再补充具体模块）：

`@.trae/rules/core/spec-index.md`

当需求涉及“安全/错误处理/测试/API 设计”或你希望更严格的约束时，使用 `GoServiceStrict` 配置加载全部模块。

使用建议：
- 涉及命令/生成链路/目录位置时，先以仓库现状为准：根目录 `makefile`、`go.mod`、现有 `internal/*` 代码与 `dep/` 目录。
- 仓库命名不作为判断依据：通过 `buf.gen.yaml`/`buf.yaml`/`buf.work.yaml` 识别微服务仓库与 Proto 仓库。
- 若规则文字与仓库现状冲突，优先遵循仓库现状并保持输出一致性。
