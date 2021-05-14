package utils

import (
	"database/sql"

	"gorm.io/gorm"
)

type TruncateTableExecutor struct {
	db *gorm.DB
}

func InitTruncateTableExecutor(db *gorm.DB) TruncateTableExecutor {
	return TruncateTableExecutor{
		db,
	}
}

func (executor *TruncateTableExecutor) TruncateTable(tableName string) {
	executor.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("TRUNCATE TABLE @tableName RESTART IDENTITY CASCADE;", sql.Named("tableName", tableName)).Error; err != nil {
			tx.Rollback()
			return err
		}

		return tx.Commit().Error
	})
}
