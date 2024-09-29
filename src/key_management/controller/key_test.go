package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/go-ja-holidays/logging"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/model"
	"github.com/kynmh69/go-ja-holidays/util"
)

func TestMain(m *testing.M) {
	setUp()
	defer tearDown()
	res := m.Run()
	os.Exit(res)
}

const ViewDir = "key_management/view/*/*.html"

func TestKeyManagement_Retrieve(t *testing.T) {
	r := gin.Default()
	util.SetUp()
	mg := NewKeyManagement("key")
	target := "/manage/key"
	r.GET(target, mg.Retrieve)
	req := httptest.NewRequest(http.MethodGet, target, nil)
	rec := httptest.NewRecorder()
	ctx := gin.CreateTestContextOnly(rec, r)
	ctx.Request = req
	wd, _ := util.JoinProjectRootPath(ViewDir)
	r.LoadHTMLGlob(wd)
	type fields struct {
		ControllerName string
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test OK",
			fields: fields{
				ControllerName: "key",
			},
			args: args{
				c: ctx,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := KeyManagement{
				ControllerName: tt.fields.ControllerName,
			}
			k.Retrieve(tt.args.c)
		})
	}
}

func TestKeyManagement_GetControllerName(t *testing.T) {
	type fields struct {
		ControllerName string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "test ok",
			fields: fields{ControllerName: "key"},
			want:   "key",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := KeyManagement{
				ControllerName: tt.fields.ControllerName,
			}
			if got := k.GetControllerName(); got != tt.want {
				t.Errorf("KeyManagement.GetControllerName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func setUp() {
	_ = os.Setenv("PSQL_HOSTNAME", "localhost")
	_ = os.Setenv("DATABASE", "unittest")
	url := "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"
	logging.LoggerInitialize()
	database.ConnectDatabase()
	util.CreateHolidayData(url)
	db := database.GetDbConnection()
	u := uuid.New()
	apikey := model.ApiKey{Key: u.String()}
	res, err := db.Insert(model.TABLE_API_KEY).Rows(
		apikey,
	).
		Executor().Exec()
	if err != nil {
		log.Fatalln(err)
	}
	r, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	log.Println(r)
}

func tearDown() {
	_ = os.Unsetenv("PSQL_HOSTNAME")
	_ = os.Unsetenv("DATABASE")
	db := database.GetDbConnection()
	if _, err := db.Delete(model.TABLE_API_KEY).Executor().Exec(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Teardown.")
}

func TestNewKeyManagement(t *testing.T) {
	name := "test handler"
	type args struct {
		controllerName string
	}
	tests := []struct {
		name string
		args args
		want *KeyManagement
	}{
		{
			name: "test ok",
			args: args{
				controllerName: name,
			},
			want: &KeyManagement{ControllerName: name},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKeyManagement(tt.args.controllerName); !reflect.DeepEqual(got.ControllerName, tt.want.ControllerName) {
				t.Errorf("NewKeyManagement() = %v, want %v", got, tt.want)
			}
		})
	}
}
