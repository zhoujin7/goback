package goback

import (
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestInstance(t *testing.T) {
	r := Instance()
	assert := assert.New(t)
	assert.NotNil(r)
	assert.NotNil(r.handlerFnMap)
	assert.NotNil(r.pathParamStore)
	assert.NotNil(r.pool)
	for reqMethod := range reqMethods {
		assert.NotNil(r.handlerFnMap[reqMethod])
		assert.NotNil(r.pathParamStore[reqMethod])
	}
}
