package main

// #cgo CFLAGS: -I./include
// #include <client.h>
// #include <stdint.h>
// #include <errno.h>
import "C"
import (
	"context"

	"github.com/pokt-network/poktroll/pkg/client"
	"github.com/pokt-network/poktroll/pkg/client/events"
)

//export NewEventsQueryClient
func NewEventsQueryClient(cometWebsocketURLCString *C.char) C.go_ref {
	// TODO_TECHDEBT: support opts args.
	cometWebsocketURL := C.GoString(cometWebsocketURLCString)
	eventsQueryClient := events.NewEventsQueryClient(cometWebsocketURL)

	return SetGoMem(eventsQueryClient)
}

//export EventsQueryClientEventsBytes
func EventsQueryClientEventsBytes(clientRef C.go_ref, query *C.char, cErr **C.char) C.go_ref {
	// TODO_CONSIDERATION: Could support a version of methods which receive a go context, created elsewhere..
	ctx := context.Background()

	eventsQueryClient, err := GetGoMem[client.EventsQueryClient](clientRef)
	if err != nil {
		*cErr = C.CString(err.Error())
		return CNilRef
	}

	eventsBytesObs, err := eventsQueryClient.EventsBytes(ctx, C.GoString(query))
	if err != nil {
		*cErr = C.CString(err.Error())
		return CNilRef
	}

	return SetGoMem(eventsBytesObs)
}
