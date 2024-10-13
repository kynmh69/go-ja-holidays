package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/logging"
	"github.com/kynmh69/go-ja-holidays/model"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("PSQL_HOSTNAME", "localhost")
	_ = os.Setenv("DATABASE", "unittest")
	logging.LoggerInitialize()
	database.ConnectDatabase()

	defer tearDown()

	db := database.GetDbConnection()
	if err := db.AutoMigrate(
		&model.ApiKey{},
		&model.HolidayData{},
	); err != nil {
		logging.GetLogger().Panicln(err)
	}
	if code := m.Run(); code > 0 {
		logging.GetLogger().Panicln("Test failed with code ", code)
	}
	logging.GetLogger().Infoln("Test passed")
}

func TestAuth(t *testing.T) {
	r := gin.Default()
	u, _ := uuid.NewUUID()
	var apiKey model.ApiKey
	apiKeyError := model.ApiKey{Key: u}
	db := database.GetDbConnection()
	db.Create(&apiKey)
	tests := []struct {
		name       string
		apiKey     model.ApiKey
		statusCode int
	}{
		{
			name:       "Test Auth",
			apiKey:     apiKey,
			statusCode: http.StatusOK,
		},
		{
			name:       "Test Auth Error",
			apiKey:     apiKeyError,
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("X-API-KEY", tt.apiKey.Key.String())
			w := httptest.NewRecorder()
			ctx := gin.CreateTestContextOnly(w, r)
			ctx.Request = req
			Auth(ctx)
			if tt.statusCode != w.Code {
				t.Errorf("expected %d but got %d", tt.statusCode, w.Code)
			}
		})
	}
}

func tearDown() {
	db := database.GetDbConnection()
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).
		Delete(&model.ApiKey{})
}
