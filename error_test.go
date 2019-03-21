package serr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	code := NewErrCode("ERR_CODE")
	serr := New(code, "hi")

	assert.Equal(t, "ERR_CODE hi", serr.Error())
	assert.Equal(t, serr.String(), serr.Error())
	assert.Equal(t, 0, serr.Len())
	assert.Nil(t, serr.ErrorOrNil())
}

func Test_NewWithData(t *testing.T) {
	code := NewErrCode("ERR_CODE")
	serr := NewWithData(code, "hi", 123)

	assert.Equal(t, "ERR_CODE hi", serr.Error())
	assert.Equal(t, serr.String(), serr.Error())
	assert.Equal(t, 0, serr.Len())
	assert.Equal(t, 123, serr.Data)
	assert.Nil(t, serr.ErrorOrNil())
}

func Test_Error_Append(t *testing.T) {
	code := NewErrCode("ERR_CODE")
	serr := New(code, "hi")

	serr.Append(errors.New("Error 1"))
	serr.Append(errors.New("Error 2"))
	serr.Append(errors.New("Error 3"))

	assert.Equal(t, "ERR_CODE hi\n0. Error 1\n1. Error 2\n2. Error 3", serr.Error())
	assert.Equal(t, serr.String(), serr.Error())
	assert.Equal(t, 3, serr.Len())
	assert.Equal(t, serr, serr.ErrorOrNil())

	serr.ErrorFormat = func([]error) string {
		return ": CUSTOM!!!"
	}
	assert.Equal(t, "ERR_CODE hi: CUSTOM!!!", serr.Error())
	assert.Equal(t, serr.String(), serr.Error())
}
