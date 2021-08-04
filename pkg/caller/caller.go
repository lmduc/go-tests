package caller

import "runtime"

func Caller(skip int) string {
	_, b, _, _ := runtime.Caller(skip)
	return b
}
