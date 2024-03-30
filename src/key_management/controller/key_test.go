package controller

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/key_management/template"
	"github.com/kynmh69/go-ja-holidays/model"
	"github.com/kynmh69/go-ja-holidays/util"
	"github.com/labstack/echo/v4"
)

func TestMain(m *testing.M) {
	setUp()
	defer tearDown()
	res := m.Run()
	os.Exit(res)
}

const VIEW_DIR = "../view/*.html"

func TestKeyManagement_Retrieve(t *testing.T) {
	e := echo.New()
	util.EchoLoggerInitialize(e)
	req := httptest.NewRequest(http.MethodGet, "/manage/key", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	wd, _ := os.Getwd()
	log.Println(wd)
	e.Renderer = template.NewTemplate(VIEW_DIR)
	type fields struct {
		ControllerName string
	}
	type args struct {
		c echo.Context
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
				c: c,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := KeyManagement{
				ControllerName: tt.fields.ControllerName,
			}
			if err := k.Retrieve(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("KeyManagement.Retrieve() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKeyManagement_Create(t *testing.T) {
	e := echo.New()
	util.EchoLoggerInitialize(e)
	req := httptest.NewRequest(http.MethodPost, "/manage/key/create", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	e.Renderer = template.NewTemplate(VIEW_DIR)
	type fields struct {
		ControllerName string
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test ok",
			fields: fields{
				ControllerName: "key",
			},
			args: args{
				c: c,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := KeyManagement{
				ControllerName: tt.fields.ControllerName,
			}
			if err := k.Create(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("KeyManagement.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKeyManagement_Update(t *testing.T) {
	e := echo.New()
	util.EchoLoggerInitialize(e)
	req := httptest.NewRequest(http.MethodPut, "/manage/key", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	e.Renderer = template.NewTemplate(VIEW_DIR)
	type fields struct {
		ControllerName string
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "test error",
			fields:  fields{ControllerName: "key"},
			args:    args{c: c},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := KeyManagement{
				ControllerName: tt.fields.ControllerName,
			}
			if err := k.Update(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("KeyManagement.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKeyManagement_Delete(t *testing.T) {
	e := echo.New()
	util.EchoLoggerInitialize(e)
	req := httptest.NewRequest(http.MethodDelete, "/manage/key", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	e.Renderer = template.NewTemplate(VIEW_DIR)
	type fields struct {
		ControllerName string
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "test ok",
			fields:  fields{ControllerName: "key"},
			args:    args{c: c},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := KeyManagement{
				ControllerName: tt.fields.ControllerName,
			}
			if err := k.Delete(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("KeyManagement.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
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
	os.Setenv("PSQL_HOSTNAME", "localhost")
	os.Setenv("DATABASE", "unittest")
	url := "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"
	database.ConnectDatabase()
	util.CreateHolidayData(url)
	db := database.GetDbConnection()
	uuid := uuid.New()
	apikey := model.ApiKey{Key: uuid.String()}
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
	os.Unsetenv("PSQL_HOSTNAME")
	os.Unsetenv("DATABASE")
	db := database.GetDbConnection()
	if _, err := db.Delete(model.TABLE_API_KEY).Executor().Exec(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Teardown.")
}

func TestNewKeyManagement(t *testing.T) {
	name := "test controller"
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
