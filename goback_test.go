package goback

import "testing"
import "github.com/stretchr/testify/assert"

func TestInstance(t *testing.T) {
	r := Instance()
	assert := assert.New(t)
	assert.NotNil(r)
}
