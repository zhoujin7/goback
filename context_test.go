package goback

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestAndResponse(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := newContext(rec, req)
	ctx.response.WriteHeader(404)
	ctx.response.Write([]byte("404 Not Found"))
	assert := assert.New(t)
	assert.Equal(req, ctx.Request())
	assert.Equal(404, ctx.Response().StatusCode())
	assert.Equal(rec, ctx.Response().ResponseWriter)
	assert.Equal("404 Not Found", rec.Body.String())
}

func TestHTML(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := newContext(rec, req)
	assert := assert.New(t)
	err := ctx.HTML(200, "<h1>hello</h1>")
	if assert.NoError(err) {
		assert.Equal(rec.Code, 200)
		assert.Equal(rec.Header().Get("Content-Type"), "text/html; charset=utf-8")
		assert.Equal(rec.Body.String(), "<h1>hello</h1>")
	}
}

func TestString(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := newContext(rec, req)
	assert := assert.New(t)
	err := ctx.String(500, "500 Internal Server Error")
	if assert.NoError(err) {
		assert.Equal(rec.Code, 500)
		assert.Equal(rec.Header().Get("Content-Type"), "text/plain; charset=utf-8")
		assert.Equal(rec.Body.String(), "500 Internal Server Error")
	}
}

func TestJSON(t *testing.T) {
	type user struct {
		Name   string `json:"name"`
		Gender string `json:"gender"`
	}
	testUser := user{"zhoujinqi", "male"}
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := newContext(rec, req)
	assert := assert.New(t)
	err := ctx.JSON(200, testUser)
	if assert.NoError(err) {
		assert.Equal(rec.Code, 200)
		assert.Equal(rec.Header().Get("Content-Type"), "application/json;charset=UTF-8")
		assert.Equal(rec.Body.String(), `{"name":"zhoujinqi","gender":"male"}`)
	}
}

func TestPathParamValue(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := newContext(rec, req)
	ctx.setPathParamValue("user", "zhoujinqi")
	ctx.setPathParamValue("user", "lisi")
	ctx.setPathParamValue("page", "1")
	assert := assert.New(t)
	assert.Equal(ctx.PathParamValue("user"), "zhoujinqi")
	assert.Equal(ctx.PathParamValue("page"), "1")
	assert.Equal(ctx.PathParamValueByIndex("user", 1), "lisi")
}
