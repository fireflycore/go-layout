package data

import (
	"context"
	"go-layout/internal/biz/repo"

	"gorm.io/gorm"
)

type transaction struct {
	data *Data
}

func NewTransaction(data *Data) repo.Transaction {
	return &transaction{data: data}
}

// ExecTx 在同一个数据库事务中执行回调，并把当前事务句柄透传给下游 Repo。
func (r *transaction) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 将事务DB放入context，让Repo方法能够获取到
		txCtx := context.WithValue(ctx, repo.TransactionContextKey, tx)
		return fn(txCtx)
	})
}
