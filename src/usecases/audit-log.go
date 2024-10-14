package usecases

import (
	"audit-log/src/domain/models"
	"audit-log/src/domain/repositories"
	"encoding/json"
	"time"
)

func CreateLog(req *models.CreateLog) (*models.AuditLog, error) {
	beforeJSON, err := json.Marshal(req.Before)
	if err != nil {
		return nil, err
	}

	afterJSON, err := json.Marshal(req.After)
	if err != nil {
		return nil, err
	}

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}
	actionTime, err := time.ParseInLocation("2006-01-02 15:04:05", req.ActionTime, location)
	if err != nil {
		return nil, err // Handle the error if parsing fails
	}

	// Create a new AuditLog instance
	log := &models.AuditLog{
		Module:     req.Module,
		ActionType: req.ActionType,
		SearchKey:  req.SearchKey,
		Before:     string(beforeJSON),
		After:      string(afterJSON),
		ActionBy:   req.ActionBy,
		ActionTime: actionTime,
	}

	// Insert the log into the database
	createdLog, err := repositories.InsertAuditLog(log)
	if err != nil {
		return nil, err // Return any error that occurs during insertion
	}

	return createdLog, nil // Return the created log
}
