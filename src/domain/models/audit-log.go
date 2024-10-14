package models

import (
	"time"
)

type AuditLog struct {
	ID         int       `gorm:"primaryKey;column:id"`
	Module     string    `gorm:"column:module"`
	ActionType string    `gorm:"column:action_type"`
	SearchKey  string    `gorm:"column:search_key"`
	Before     string    `gorm:"column:before_data"`
	After      string    `gorm:"column:after_data"`
	ActionBy   string    `gorm:"column:action_by"`
	ActionTime time.Time `gorm:"column:action_time"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (AuditLog) TableName() string {
	return "audit_log"
}

type LogResponse struct {
	Module     string `json:"module"`
	ActionType string `json:"action_type"`
	SearchKey  string `json:"search_key"`
	Before     string `json:"before_data"`
	After      string `json:"after_data"`
	ActionBy   string `json:"action_by"`
	ActionTime string `json:"action_time"`
}

type CreateLog struct {
	Module     string      `json:"module" validate:"required"`
	ActionType string      `json:"action_type" validate:"required"`
	SearchKey  string      `json:"search_key" validate:"required"`
	Before     interface{} `json:"before_data" validate:"required"`
	After      interface{} `json:"after_data" validate:"required"`
	ActionBy   string      `json:"action_by" validate:"required"`
	ActionTime string      `json:"action_time" validate:"required"`
}
