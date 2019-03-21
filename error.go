package supererr

type Error interface {
	// Provided for compatibility with built-in error type
	Error() string
	// Custom error code in the format
	ErrCode() ErrCode
	// Returns the place where the error occurred
	Origin() WrapPlace
	// List of places where the error passed by
	WrapedBy() []WrapPlace
	// Appends the place that called this function to the WrapedBy()
	Wrap()
	// Multi-line string showing the full execution stack
	FullStack() string
	// Error specific information
	Data() interface{}
}

type superErr struct {
	message  string
	code     ErrCode
	origin   WrapPlace
	passedBy []WrapPlace
	stack    string
	data     interface{}
}

func New(code ErrCode, msg string) *Error {
	return nil
}

func NewWithData(code ErrCode, msg string, data interface{}) *Error {
	return nil
}

func (this *superErr) Error() string {
	return this.message
}

func (this *superErr) ErrCode() ErrCode {
	return this.code
}

func (this *superErr) Origin() WrapPlace {
	return this.origin
}

func (this *superErr) PassedBy() []WrapPlace {
	return this.passedBy
}

func (this *superErr) FullStack() string {
	return this.stack
}

func (this *superErr) Data() interface{} {
	return this.data
}
