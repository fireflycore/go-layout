package repo

import "context"

const TransactionContextKey = "TX_DB"

// Transaction 定义跨多个 Repo 写入时可选接入的事务抽象。
// 当业务需要原子写入时，再由 UseCase 显式依赖并接入。
type Transaction interface {
	ExecTx(ctx context.Context, fn func(ctx context.Context) error) error
}
