# Trae 规则（go-layout）

本目录用于约束 AI 的代码输出与本项目保持一致：分层边界、Proto First、生成链路（buf/goverter/wire）、日志与错误处理风格。

推荐在对话中仅引用入口文件（需要更严格时再补充具体模块）：

`@.trae/rules/core/spec-index.md`

当需求涉及“安全/错误处理/测试”或你希望更严格的约束时，使用 `GoServiceStrict` 配置加载全部模块。
