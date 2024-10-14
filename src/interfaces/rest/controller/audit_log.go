package controller

import (
	"net/http"
	"time"

	"audit-log/src/domain/models"
	"audit-log/src/usecases"

	"github.com/labstack/echo/v4"
)

func CreateLog(c echo.Context) error {
	var req models.CreateLog

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"data":    nil,
			"message": "invalid request body",
			"error":   err.Error(),
		})
	}

	createdLog, err := usecases.CreateLog(&req)
	if err != nil {
		return c.JSON(echo.ErrBadRequest.Code, echo.Map{
			"data":    nil,
			"message": err.Error(),
		})
	}

	LogResponse := models.LogResponse{
		Module:     createdLog.Module,
		ActionType: createdLog.ActionType,
		SearchKey:  createdLog.SearchKey,
		Before:     createdLog.Before,
		After:      createdLog.After,
		ActionBy:   createdLog.ActionBy,
		ActionTime: createdLog.ActionTime.In(time.FixedZone("WIB", 7*3600)).Format("2006-01-02 15:04:05"),
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"data":    LogResponse,
		"message": "Log created successfully",
	})
}
