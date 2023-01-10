package repository

import (
	"api_client/usecase/repository"
	"context"
	"fmt"

	"gorm.io/gorm"
)

var txKey = struct{}{}

type txRepository struct {
	db *gorm.DB
}

func (tr *txRepository) GetDBConn() *gorm.DB {
	return tr.db
}

func NewTxRepository(db *gorm.DB) repository.TxRepository {
	return &txRepository{db}
}

// DoInTx は，トランザクションオブジェクトを生成し，contextに入れて，次の関数を実行し，errorに応じて適切にrollbackやcommitを行います．
func (tr *txRepository) DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	// txを生成する
	conn := tr.GetDBConn()
	tx := conn.Begin()

	// txをcontextに入れて次の関数を実行する
	ctx = context.WithValue(ctx, &txKey, tx)
	v, err := f(ctx)
	if err != nil {
		_ = tx.Rollback()
		return v, fmt.Errorf("rollback: %w", err)
	}
	if result := tx.Commit(); result.Error != nil {
		_ = tx.Rollback()
		return v, fmt.Errorf("failed to commit: rollback: %v", result.Error)
	}
	return v, nil
}

// context.Contextからトランザクションを取得する
func GetTx(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(&txKey).(*gorm.DB)
	return tx, ok
}
