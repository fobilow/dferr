package dferr

type ErrorFunc func() error
type ErrorHandler func(err error)
func voidErrorHandler(error) {}

func This(f ErrorFunc) {
	voidErrorHandler(f())
}

func ThisWithHandler(f ErrorFunc, handler ErrorHandler) {
	handler(f())
}