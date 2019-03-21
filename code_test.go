package serr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewErrCode_1(t *testing.T) {
	vals := []string{"hi"}
	code := NewErrCode(vals...)
	assert.Equal(t, vals, code.codes)
}

func Test_NewErrCode_2(t *testing.T) {
	vals := []string{"hi", "sfd", "efd", "efwwr"}
	code := NewErrCode(vals...)
	assert.Equal(t, vals, code.codes)
}

func Test_SameErrCode_1(t *testing.T) {
	code1 := NewErrCode("ERR_GENERIC")
	code2 := NewErrCode("ERR_GENERIC")
	assert.True(t, SameErrCode(code1, code2))
	assert.True(t, code1.Equal(code2))
	assert.True(t, code2.Equal(code1))
}

func Test_SameErrCode_2(t *testing.T) {
	code1 := NewErrCode("ERR_SPECIFIC", "ERR_GENERIC")
	code2 := NewErrCode("ERR_GENERIC")
	code3 := NewErrCode("ERR_SPECIFIC")

	assert.True(t, SameErrCode(code1, code2))
	assert.True(t, code1.Equal(code2))
	assert.True(t, code2.Equal(code1))

	assert.True(t, SameErrCode(code1, code3))
	assert.True(t, code1.Equal(code3))
	assert.True(t, code3.Equal(code1))

	assert.False(t, SameErrCode(code3, code2))
	assert.False(t, code3.Equal(code2))
	assert.False(t, code2.Equal(code3))
}

func Test_Code_String(t *testing.T) {
	code1 := NewErrCode("ERR_SPECIFIC", "ERR_GENERIC")
	code2 := NewErrCode("ERR_GENERIC")
	assert.Equal(t, code1.String(), "ERR_SPECIFIC")
	assert.Equal(t, code2.String(), "ERR_GENERIC")
}
