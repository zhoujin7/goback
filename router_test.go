package goback

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter_Get(t *testing.T) {
	assert := assert.New(t)
	r := Instance()
	r.Get("/", func(ctx *Context) error {
		return ctx.String(200, "hello")
	})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	assert.Equal(rec.Code, 200)
	assert.Equal(rec.Header().Get("Content-Type"), "text/plain; charset=utf-8")
	assert.Equal(rec.Body.String(), "hello")
}

func TestRouter_Delete(t *testing.T) {
	assert := assert.New(t)
	r := Instance()
	r.Delete("/user/:userId", func(ctx *Context) error {
		return ctx.String(200, "delete "+ctx.PathParamValue("userId"))
	})
	req := httptest.NewRequest(http.MethodDelete, "/user/1", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	assert.Equal(rec.Code, 200)
	assert.Equal(rec.Header().Get("Content-Type"), "text/plain; charset=utf-8")
	assert.Equal(rec.Body.String(), "delete 1")
}

func TestRouter_Use(t *testing.T) {
	assert := assert.New(t)
	r := Instance()
	r.Use(func(fn HandlerFn) HandlerFn {
		return func(ctx *Context) error {
			return nil
		}
	})
	assert.NotNil(r.middlewareChain)
	assert.Len(r.middlewareChain, 1)
}
