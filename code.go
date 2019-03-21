package serr

type ErrCode struct {
	codes []string
}

// Creates a new error code, the first value will be the "main" code and the others will be more generic or alternative versions of the error. Example: NewErrCode("TCP_SOCKET_TIME_OUT", "SOCKET_FAIL", "SOCKET_TIME_OUT", "TIME_OUT", "NETWORK_FAIL")
func NewErrCode(codes ...string) *ErrCode {
	ans := new(ErrCode)
	ans.codes = make([]string, len(codes))
	for i := 0; i < len(codes); i++ {
		ans.codes[i] = codes[i]
	}
	return ans
}

func SameErrCode(code1, code2 *ErrCode) bool {
	for i := 0; i < len(code1.codes); i++ {
		for j := 0; j < len(code2.codes); j++ {
			if code1.codes[i] == code2.codes[j] {
				return true
			}
		}
	}
	return false
}

func (this *ErrCode) Equal(other *ErrCode) bool {
	return SameErrCode(this, other)
}

func (this *ErrCode) String() string {
	return this.codes[0]
}
