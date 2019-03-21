package serr

import (
	"fmt"
	"runtime"
)

type WrapPlace struct {
	ok       bool
	line     int
	function string
	file     string
}

func (this WrapPlace) String() string {
	if !this.ok {
		return fmt.Sprintf("[unknown place]")
	}
	if this.function == "" {
		return fmt.Sprintf("%s:%d", this.file, this.line)
	}
	return fmt.Sprintf("%s:%d", this.function, this.line)
}

func (this WrapPlace) Function() string {
	return this.function
}

func (this WrapPlace) Line() int {
	return this.line
}

func newWrap(n int) WrapPlace {
	ans := WrapPlace{}
	var pc uintptr

	pc, ans.file, ans.line, ans.ok = runtime.Caller(n + 1)
	if !ans.ok {
		return ans
	}
	function := runtime.FuncForPC(pc)
	if function == nil {
		return ans
	}
	ans.function = function.Name()
	return ans
}

// Returns a string with the execution stack for this goroutine
func fullStack() string {
	buf := make([]byte, 1000000)
	runtime.Stack(buf, false)
	return string(buf)
}
