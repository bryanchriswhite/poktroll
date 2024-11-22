package main

// #cgo CFLAGS: -I./include
// #include <events_query_client.h>
import "C"

var errorCodeMessageMap = map[int]string{
	0: "EventsBytesSyncError",
	1: "EventsBytesAsyncError",
}

//export GetErrMessage
func GetErrMessage(code int) string {
	return errorCodeMessageMap[code]
}
