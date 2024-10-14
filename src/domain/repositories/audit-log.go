package repositories

import (
	"audit-log/src/domain/models"
	"audit-log/src/infrastructure/database/mysql"
)

func InsertAuditLog(auditLog *models.AuditLog) (*models.AuditLog, error) {
	if err := mysql.GetMySQLWriteDB().Create(auditLog).Error; err != nil {
		return nil, err
	}
	return auditLog, nil
}
