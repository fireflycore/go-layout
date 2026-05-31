package data

import (
	"context"
	"go-layout/internal/biz/repo"

	"gorm.io/gorm"
)

// transaction 负责为跨 Repo 写入流程提供统一事务执行入口。
type transaction struct {
	data *Data
}

// NewTransaction 创建事务执行器。
func NewTransaction(data *Data) repo.Transaction {
	return &transaction{data: data}
}

// ExecTx 在同一个数据库事务中执行回调，并把当前事务句柄透传给下游 Repo。
func (r *transaction) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 把当前事务句柄写入 context，供下游 Repo 统一复用。
		txCtx := context.WithValue(ctx, repo.TransactionContextKey, tx)
		return fn(txCtx)
	})
}
