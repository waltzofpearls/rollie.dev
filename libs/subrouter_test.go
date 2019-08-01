package libs_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/waltzofpearls/rollie.dev/libs"
)

func TestRedirectHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()

	sr := &libs.Subrouter{}
	handler := libs.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		return sr.RedirectHandler(w, r, "")
	})
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
}

func TestHtmlResponseHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()

	tpl := libs.NewTemplate("../")
	sr := &libs.Subrouter{Template: tpl}
	handler := libs.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		return sr.HtmlResponseHandler(w, r, "test", map[string]interface{}{})
	})
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestJsonResponseHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()

	sr := &libs.Subrouter{}
	handler := libs.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		return sr.JsonResponseHandler(w, r, struct{}{})
	})
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestJsonNotFoundHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()

	sr := &libs.Subrouter{}
	handler := libs.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		return sr.JsonNotFoundHandler(w, r)
	})
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestJsonErrorHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()

	sr := &libs.Subrouter{}
	handler := libs.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		return sr.JsonErrorHandler(w, r, errors.New(""))
	})
	handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
