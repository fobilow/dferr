package dferr

type ErrorFunc func() error
type ErrorHandler func(err error)

func voidErrorHandler(error) {}

func Func(f ErrorFunc) {
	voidErrorHandler(f())
}

func FuncWithHandler(f ErrorFunc, handler ErrorHandler) {
	handler(f())
}
