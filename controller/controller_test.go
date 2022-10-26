package controller_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mudi25/pretest/controller"
)

func TestMain(m *testing.M) {
	m.Run()
}
func TestNewController(t *testing.T) {
	e := controller.NewController()
	if e == nil {
		t.Error("expect controller not nil")
	}
}
func Language() *bytes.Buffer {
	body, _ := json.Marshal(map[string]interface{}{
		"language":        "C",
		"appeared":        1972,
		"created":         []string{"Dennis Ritchie"},
		"functional":      true,
		"object-oriented": false,
		"relation": map[string]interface{}{
			"influenced-by": []string{
				"B",
				"ALGOL 68",
				"Assembly",
				"FORTRAN",
			},
			"influences": []string{
				"C++",
				"Objective-C",
				"C#",
				"Java",
				"JavaScript",
				"PHP",
				"Go",
			},
		},
	})
	return bytes.NewBuffer(body)
}
func CreateTest(method string, url string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, body)
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	return ctx, rec
}
func TestGetHello(t *testing.T) {
	ctx, rec := CreateTest("GET", "/", nil)
	controller.GetHello(ctx)
	if rec.Code != 200 {
		t.Error("expect status 200")
	}
}
func TestGetLanguage(t *testing.T) {
	ctx, rec := CreateTest("GET", "/", nil)
	controller.GetLanguage(ctx)
	if rec.Code != 200 {
		t.Error("expect status 200")
	}
}
func TestGetPalindrom(t *testing.T) {
	t.Run("BadRequest", func(t *testing.T) {
		ctx, rec := CreateTest("GET", "/", nil)
		controller.GetPalindrom(ctx)
		if rec.Code != 400 {
			t.Error("expect status 400")
		}
	})
	t.Run("Success", func(t *testing.T) {
		body, _ := json.Marshal(map[string]interface{}{
			"text": "101",
		})
		ctx, rec := CreateTest("GET", "/", bytes.NewBuffer(body))
		controller.GetPalindrom(ctx)
		if rec.Code != 200 {
			t.Error("expect status 200")
		}
	})
}
func TestCreateLanguage(t *testing.T) {
	t.Run("BadRequest", func(t *testing.T) {
		ctx, rec := CreateTest("POST", "/", nil)
		controller.CreateLanguage(ctx)
		if rec.Code != 400 {
			t.Error("expect status 400")
		}
	})
	t.Run("Success", func(t *testing.T) {
		ctx, rec := CreateTest("POST", "/", Language())
		controller.CreateLanguage(ctx)
		if rec.Code != 200 {
			t.Error("expect status 200")
		}
	})
}
func TestGetLanguages(t *testing.T) {
	ctx, rec := CreateTest("Get", "/", nil)
	controller.GetLanguages(ctx)
	if rec.Code != 200 {
		t.Error("expect status 200")
	}
}
func TestGetLanguageById(t *testing.T) {
	t.Run("BadRequest", func(t *testing.T) {
		ctx, rec := CreateTest("GET", "/a", nil)
		controller.GetLanguageById(ctx)
		if rec.Code != 400 {
			t.Error("expect status 400")
		}
	})
}
func TestPatchLanguage(t *testing.T) {
	t.Run("BadRequest", func(t *testing.T) {
		ctx, rec := CreateTest("PATCH", "/a", nil)
		controller.PatchLanguage(ctx)
		if rec.Code != 400 {
			t.Error("expect status 400")
		}
	})
}
func TestDeleteLanguage(t *testing.T) {
	t.Run("BadRequest", func(t *testing.T) {
		ctx, rec := CreateTest("DELETE", "/a", nil)
		controller.DeleteLanguage(ctx)
		if rec.Code != 400 {
			t.Error("expect status 400")
		}
	})
}
