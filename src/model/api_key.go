package model

import (
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/labstack/echo/v4"
)

const TABLE_API_KEY = "api_key"
const COLUMN_KEY = "key"
const COLUMN_CREATED_AT = "created_at"

type ApiKey struct {
	Id        string    `db:"id" goqu:"skipinsert"`
	Key       string    `db:"key"`
	CreatedAt time.Time `db:"created_at" goqu:"skipinsert"`
}

func GetApiKeys() ([]ApiKey, error) {
	var apiKeys []ApiKey
	db := database.GetDbConnection()
	err := db.From(TABLE_API_KEY).ScanStructs(&apiKeys)
	return apiKeys, err
}

func CreateApiKey(c echo.Context) ([]ApiKey, error) {
	var apiKeys []ApiKey
	logger := c.Logger()
	key := uuid.New()
	db := database.GetDbConnection()
	result, err := db.Insert(TABLE_API_KEY).
		Rows(
			ApiKey{Key: key.String()},
		).
		Executor().Exec()
	if err != nil {
		return apiKeys, err
	}

	id, _ := result.RowsAffected()
	logger.Info("Create API Key.", id)
	apiKeys, err = GetApiKeys()
	return apiKeys, err
}

func DeleteApiKey(c echo.Context) ([]ApiKey, error) {
	logger := c.Logger()
	db := database.GetDbConnection()

	anHourAgo := time.Now().Add(-1 * time.Hour)

	result, err := db.Delete(TABLE_API_KEY).Where(
		goqu.C(COLUMN_CREATED_AT).Lt(anHourAgo),
	).Executor().Exec()

	if err != nil {
		return nil, err
	}
	row, _ := result.RowsAffected()
	logger.Debug("Row affected", row)
	apiKeys, err := GetApiKeys()
	return apiKeys, err
}
