package serr

import "fmt"

type ErrorFormatFunc func([]error) string

type SuperError struct {
	Message     string
	Code        *ErrCode
	Origin      WrapPlace
	PassedBy    []WrapPlace
	Stack       string
	Errors      []error
	ErrorFormat ErrorFormatFunc
	// Application specific data or context for the error
	Data interface{}
}

func New(code *ErrCode, msg string) *SuperError {
	ans := new(SuperError)
	ans.Code = code
	ans.Message = msg
	ans.Origin = newWrap(1)
	ans.Stack = fullStack()
	return ans
}

func NewWithData(code *ErrCode, msg string, data interface{}) *SuperError {
	ans := new(SuperError)
	ans.Code = code
	ans.Message = msg
	ans.Origin = newWrap(1)
	ans.Stack = fullStack()
	ans.Data = data
	return ans
}

func (this *SuperError) String() string {
	return this.Error()
}

func (this *SuperError) Error() string {
	ans := this.Origin.String()
	if ans != "" {
		ans += " "
	}
	ans = this.Code.String()
	if ans != "" {
		ans += " "
	}
	ans += this.Message

	// Let the user have their special formatter functions
	if this.ErrorFormat != nil {
		ans += this.ErrorFormat(this.Errors)
		return ans
	}

	// Default formatter
	for i, sub := range this.Errors {
		ans += fmt.Sprintf("\n%d. %s", i, sub.Error())
	}
	return ans
}

// Returns the amount of sub errors
func (this *SuperError) Len() int {
	return len(this.Errors)
}

// Appends sub errors.
func (this *SuperError) Append(errs ...error) {
	if this.Errors == nil {
		this.Errors = make([]error, 0)
	}
	this.Errors = append(this.Errors, errs...)
}

// Returns nil if and only if there are no sub errors. Otherwise returns itself.
func (this *SuperError) ErrorOrNil() *SuperError {
	if this.Len() == 0 {
		return nil
	}
	return this
}
