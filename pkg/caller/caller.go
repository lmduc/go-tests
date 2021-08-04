package caller

import (
	"os"
)

func Caller(skip int) string {
	path, _ := os.Getwd()
	return path
}
