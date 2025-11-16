package repo

import "context"

const TransactionContextKey = "TX_DB"

// Transaction 事务接口
type Transaction interface {
	ExecTx(ctx context.Context, fn func(ctx context.Context) error) error
}
