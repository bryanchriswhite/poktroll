package main

// TODO_TECHDEBT: improve error handling:
//   - be conventional (error codes?)

// #cgo CFLAGS: -I./include
// #include <client.h>
import "C"

var errorCodeMessageMap = map[int]string{
	0: "EventsBytesSyncError",
	1: "EventsBytesAsyncError",
}

//export GetErrMessage
func GetErrMessage(code int) string {
	return errorCodeMessageMap[code]
}
