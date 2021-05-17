package utils

import (
	"fmt"

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

func (executor *TruncateTableExecutor) TruncateTable(tableNames []string) {
	executor.db.Transaction(func(tx *gorm.DB) error {
		for _, tableName := range tableNames {
			query := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE;", tableName)
			if err := tx.Exec(query).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

		return tx.Commit().Error
	})
}
