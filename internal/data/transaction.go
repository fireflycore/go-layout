package data

import (
	"context"
	"go-layout/internal/biz/repo"
	"gorm.io/gorm"
)

type transaction struct {
	db *gorm.DB
}

func NewTransaction(db *gorm.DB) repo.Transaction {
	return &transaction{db: db}
}

func (t *transaction) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 将事务DB放入context，让Repo方法能够获取到
		txCtx := context.WithValue(ctx, repo.TransactionContextKey, tx)
		return fn(txCtx)
	})
}
