package server

import "runtime"

var (
	// GitCommit 在构建阶段通过 ldflags 注入；本地开发默认使用 dev。
	GitCommit = "dev"
	// BuildTime 在构建阶段通过 ldflags 注入；本地开发默认使用 unknown。
	BuildTime = "unknown"
)

// buildInfo 返回管理端口对外暴露的稳定构建信息。
func buildInfo() map[string]any {
	return map[string]any{
		"git_commit": GitCommit,
		"build_time": BuildTime,
		"go_version": runtime.Version(),
	}
}
