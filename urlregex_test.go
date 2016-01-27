package urlregex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPattern(t *testing.T) {
	pattern := Pattern("some/:cool/pattern/:value")
	assert.NotNil(t, pattern)
	assert.Equal(t, pattern.Regex.String(), "^some\\/(?P<cool>.[^\\/]*)\\/pattern\\/(?P<value>.[^\\/]*)$")

	pattern = Pattern("some/:cool/pattern/:value/")
	assert.NotNil(t, pattern)
	assert.Equal(t, pattern.Regex.String(), "^some\\/(?P<cool>.[^\\/]*)\\/pattern\\/(?P<value>.[^\\/]*)\\/$")

	pattern = Pattern("/some/:cool/pattern/:value")
	assert.NotNil(t, pattern)
	assert.Equal(t, pattern.Regex.String(), "^\\/some\\/(?P<cool>.[^\\/]*)\\/pattern\\/(?P<value>.[^\\/]*)$")

	pattern = Pattern("/some/:cool/pattern/:value/")
	assert.NotNil(t, pattern)
	assert.Equal(t, pattern.Regex.String(), "^\\/some\\/(?P<cool>.[^\\/]*)\\/pattern\\/(?P<value>.[^\\/]*)\\/$")

	pattern = Pattern("/some/*")
	assert.NotNil(t, pattern)
	assert.Equal(t, pattern.Regex.String(), "^\\/some\\/.*$")

	pattern = Pattern("*")
	assert.NotNil(t, pattern)
	assert.Equal(t, pattern.Regex.String(), "^.*$")

	//named wildcard
	pattern = Pattern("/*key")
	assert.NotNil(t, pattern)
	assert.Equal(t, pattern.Regex.String(), "^\\/(?P<key>.*)$")
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

	// test wildcard
	pattern = Pattern("/blabla/*/some")
	res, err = pattern.Match("/blabla/route/some")
	assert.Nil(t, err)

	res, err = pattern.Match("/blabl/route/some")
	assert.NotNil(t, err)

	res, err = pattern.Match("/blabla/route/")
	assert.NotNil(t, err)

	res, err = pattern.Match("/blabla/route")
	assert.NotNil(t, err)

	pattern = Pattern("/blabla/*")
	res, err = pattern.Match("/blabla/route/some")
	assert.Nil(t, err)

	res, err = pattern.Match("/blabla/route/some/")
	assert.Nil(t, err)

	res, err = pattern.Match("/blabla")
	assert.NotNil(t, err)

	pattern = Pattern("*")
	res, err = pattern.Match("/some/route")
	assert.Nil(t, err)

	res, err = pattern.Match("/blabla/route/some/")
	assert.Nil(t, err)

	// combine wildcard and url-params
	pattern = Pattern("some/:cool/*/:value")
	res, err = pattern.Match("some/123/pattern/456")
	assert.Nil(t, err)
	assert.Equal(t, res["cool"], "123")
	assert.Equal(t, res["value"], "456")

	// named wildcard
	pattern = Pattern("some/*key")
	res, err = pattern.Match("some/this/is/a/key")
	assert.Nil(t, err)
	assert.Equal(t, res["key"], "this/is/a/key")
}
