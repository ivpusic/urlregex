package urlreg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPattern(t *testing.T) {
	pattern := Pattern("some/:cool/pattern/:value")
	assert.NotNil(t, pattern)
	assert.Equal(t, pattern.regex.String(), "^some\\/(?P<cool>.[^\\/]*)\\/pattern\\/(?P<value>.[^\\/]*)$")

	pattern = Pattern("some/:cool/pattern/:value/")
	assert.NotNil(t, pattern)
	assert.Equal(t, pattern.regex.String(), "^some\\/(?P<cool>.[^\\/]*)\\/pattern\\/(?P<value>.[^\\/]*)\\/$")

	pattern = Pattern("/some/:cool/pattern/:value")
	assert.NotNil(t, pattern)
	assert.Equal(t, pattern.regex.String(), "^\\/some\\/(?P<cool>.[^\\/]*)\\/pattern\\/(?P<value>.[^\\/]*)$")

	pattern = Pattern("/some/:cool/pattern/:value/")
	assert.NotNil(t, pattern)
	assert.Equal(t, pattern.regex.String(), "^\\/some\\/(?P<cool>.[^\\/]*)\\/pattern\\/(?P<value>.[^\\/]*)\\/$")
}

func TestMatch(t *testing.T) {
	pattern := Pattern("some/:cool/pattern/:value")
	res, err := pattern.Match("some/123/pattern/456")
	assert.Nil(t, err)
	assert.NotNil(t, res["cool"])
	assert.NotNil(t, res["value"])

	pattern = Pattern("some/:cool/pattern/:value/")
	res, err = pattern.Match("some/123/pattern/456/")
	assert.Nil(t, err)
	assert.NotNil(t, res["cool"])
	assert.NotNil(t, res["value"])

	pattern = Pattern("/some/:cool/pattern/:value")
	res, err = pattern.Match("/some/123/pattern/456")
	assert.Nil(t, err)
	assert.NotNil(t, res["cool"])
	assert.NotNil(t, res["value"])

	pattern = Pattern("/some/:cool/pattern/:value/")
	res, err = pattern.Match("/some/123/pattern/456/")
	assert.Nil(t, err)
	assert.NotNil(t, res["cool"])
	assert.NotNil(t, res["value"])

	// bellow should not match
	pattern = Pattern("some/:cool/pattern/:value")
	res, err = pattern.Match("some/123/pattern/456/")
	assert.NotNil(t, err)

	pattern = Pattern("some/:cool/pattern/:value/")
	res, err = pattern.Match("/some/123/pattern/456")
	assert.NotNil(t, err)

	pattern = Pattern("/some/:cool/pattern/:value")
	res, err = pattern.Match("/some/123/pattern/456/")
	assert.NotNil(t, err)

	pattern = Pattern("/some/:cool/pattern/:value/")
	res, err = pattern.Match("/some/123/pattern/456")
	assert.NotNil(t, err)

	pattern = Pattern("/some/:cool/pattern/:value/")
	res, err = pattern.Match("some/123/pattern/456/")
	assert.NotNil(t, err)
}
