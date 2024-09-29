package model

import (
	"github.com/google/uuid"
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/logging"
	"gorm.io/gorm"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	err := setUp()
	if err != nil {
		panic(err)
	}
	code := m.Run()
	defer tearDown()
	if code != 0 {
		panic(code)
	}
}

func tearDown() {

}

func setUp() error {
	_ = os.Setenv("PSQL_HOSTNAME", "localhost")
	_ = os.Setenv("PSQL_DATABASE", "unittest")
	logging.LoggerInitialize()
	database.ConnectDatabase()
	db := database.GetDbConnection()
	err := db.AutoMigrate(&ApiKey{}, &HolidayData{})
	if err != nil {
		return err
	}
	return nil
}

func TestCreateApiKey(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Create API Key",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateApiKey(); (err != nil) != tt.wantErr {
				t.Errorf("CreateApiKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteApiKey(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Delete API Key",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteApiKey(); (err != nil) != tt.wantErr {
				t.Errorf("DeleteApiKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetApiKeys(t *testing.T) {
	// Create
	id, _ := uuid.Parse("602d8dce-cff6-452c-8363-fa528aff5f3f")
	key, _ := uuid.Parse("f88f4587-e49e-4574-beb3-e72f8d050efb")
	apiKey := ApiKey{
		Id:        id,
		Key:       key,
		CreatedAt: time.Now(),
	}
	db := database.GetDbConnection()
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).
		Unscoped().
		Delete(&ApiKey{})
	db.Create(&apiKey)
	tests := []struct {
		name    string
		want    []ApiKey
		wantErr bool
	}{
		{
			name: "Get API Keys",
			want: []ApiKey{
				apiKey,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetApiKeys()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApiKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("GetApiKeys() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
