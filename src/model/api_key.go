package model

import (
	"github.com/kynmh69/go-ja-holidays/logging"
	"time"

	"github.com/google/uuid"
	"github.com/kynmh69/go-ja-holidays/database"
)

const ColumnCreatedAt = "created_at"

type ApiKey struct {
	Id        uuid.UUID `gorm:"id;primary_key;default:gen_random_uuid();"`
	Key       uuid.UUID `gorm:"api_key;not null;unique;default:gen_random_uuid();"`
	CreatedAt time.Time `gorm:"created_at"`
}

func GetApiKeys() ([]ApiKey, error) {
	var apiKeys []ApiKey
	db := database.GetDbConnection()
	err := db.Find(&apiKeys).Order(ColumnCreatedAt).Error
	return apiKeys, err
}

func GetApiKey(apiKeyStr string) (ApiKey, error) {
	logger := logging.GetLogger()
	var apiKey ApiKey
	db := database.GetDbConnection()
	err := db.Where("key = ?", apiKeyStr).First(&apiKey).Error
	if err != nil {
		logger.Debug("API key is invalid. ")
	}
	return apiKey, err
}

func CreateApiKey() error {
	logger := logging.GetLogger()
	db := database.GetDbConnection()
	var key ApiKey
	tx := db.Create(&key)
	err := tx.Error
	if err != nil {
		return err
	}
	logger.Info("Create API Key.", key.Id)
	return err
}

func DeleteApiKey() error {
	logger := logging.GetLogger()
	db := database.GetDbConnection()
	defaultLocation := time.Local
	logger.Debug(defaultLocation)
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}
	time.Local = loc
	anHourAgo := time.Now().Add(-1 * time.Hour)
	logger.Debug("an hour ago: ", anHourAgo)
	result := db.Where(ColumnCreatedAt+" <= ?", anHourAgo).Delete(&ApiKey{})
	if result.Error != nil {
		return result.Error
	}
	row := result.RowsAffected
	logger.Debug("Row affected", row)
	return nil
}
