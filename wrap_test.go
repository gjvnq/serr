package serr

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newWrap_1(t *testing.T) {
	w := newWrap(0)
	assert.True(t, strings.HasSuffix(w.file, "wrap_test.go"))
	assert.NotEqual(t, 0, w.line)
	assert.Equal(t, w.line, w.Line())
	assert.Equal(t, w.function, w.Function())
	assert.True(t, strings.HasPrefix(w.String(), "github.com/gjvnq/serr"))
}

func Test_newWrap_2(t *testing.T) {
	w := newWrap(999)
	assert.Equal(t, 0, w.line)
	assert.Equal(t, w.line, w.Line())
	assert.Equal(t, w.function, w.Function())
	assert.Equal(t, "[unknown place]", w.String())
}
