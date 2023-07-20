package goroutineid

import (
	"bytes"
	"runtime"
	"strconv"
)

func GetCurrentGoId() (uint64, error) {
	b := make([]byte, 1024)
	runtime.Stack(b, false)
	return strconv.ParseUint(string(bytes.Fields(b)[1]), 10, 64)
}
