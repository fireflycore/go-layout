package repo

import "context"

// TransactionContextKey 是把事务句柄透传给下游 Repo 时使用的 context key。
const TransactionContextKey = "TX_DB"

// Transaction 定义跨多个 Repo 写入时可选接入的事务抽象。
// 当业务需要原子写入时，再由 UseCase 显式依赖并接入。
type Transaction interface {
	// ExecTx 在同一个数据库事务中执行回调，并把当前事务句柄透传给下游 Repo。
	ExecTx(ctx context.Context, fn func(ctx context.Context) error) error
}
