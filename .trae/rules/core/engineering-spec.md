---
trigger: manual
---

# 工程行为规范（go-layout）v1.0

## [规则 1] 危险操作确认机制 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: All
说明：
- 执行高风险操作前必须获得明确确认（文件删除/批量修改、Git 提交/重置、系统配置/权限变更、数据库结构变更等）
- 确认格式必须包含：操作类型、影响范围、风险评估
- 除非用户明确要求，否则不主动执行 git commit/push 等版本控制操作

## [规则 2] 命令与路径规范 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: All
说明：
- 路径处理：始终使用双引号包裹文件路径；优先使用正斜杠 `/`；检查跨平台兼容性
- 工具优先级：优先使用 ripgrep (`rg`) > `grep`；优先使用专用工具 (Read/Write) > 系统命令
- 批量操作优先调用批量工具以提高效率

## [规则 3] 核心编程原则 (SOLID/KISS/DRY) [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- KISS (Keep It Simple, Stupid)：追求极致简洁，拒绝不必要的复杂性，优先选择直观方案
- YAGNI (You Aren't Gonna Need It)：仅实现当前明确功能，抵制过度设计，删除无用代码
- DRY (Don't Repeat Yourself)：自动识别重复模式，主动抽象复用
- SOLID：
  - SRP：单一职责
  - OCP：对扩展开放，对修改关闭
  - LSP：子类可替换父类
  - ISP：接口隔离，避免胖接口
  - DIP：依赖抽象

## [规则 4] 持续问题解决与验证 [ENABLED]
STATUS: ENABLED
PRIORITY: MEDIUM
LANGUAGE: All
说明：
- 持续工作直到问题解决，基于事实而非猜测
- 操作前充分规划，先读后写，理解现有代码再修改
- 每次代码变更都应体现编程原则，并验证修改的正确性（具体验证标准见 `quality/testing-spec.md`）
