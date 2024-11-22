package main

// #cgo CFLAGS: -I./include
// #include <events_query_client.h>
// #include <stdint.h>
// #include <errno.h>
import "C"
import (
	"context"

	"github.com/pokt-network/poktroll/pkg/client"
	"github.com/pokt-network/poktroll/pkg/client/events"
)

var (
	goMemoryMap = map[int64]any{}
	nextGoMemID = int64(0)
)

// main is a dummy function to satisfy the cgo requirements.
func main() {}

func SetGoMemory(value any) int64 {
	nextGoMemID++
	goMemoryMap[nextGoMemID] = value
	return nextGoMemID
}

func GetGoMemory(id int64) any {
	return goMemoryMap[id]
}

func FreeGoMemory(id int64) {
	delete(goMemoryMap, id)
}

//export NewEventsQueryClient
func NewEventsQueryClient(cometWebsocketURLCString *C.char) C.int64_t {
	// TODO_TECHDEBT: support opts args.
	cometWebsocketURL := C.GoString(cometWebsocketURLCString)
	eventsQueryClient := events.NewEventsQueryClient(cometWebsocketURL)

	return C.int64_t(SetGoMemory(eventsQueryClient))
}

//export EventsQueryClientEventsBytes
func EventsQueryClientEventsBytes(clientID int64, query *C.char, cErr **C.char) C.int64_t {
	// TODO_CONSIDERATION: Could support a version of methods which receive a go context, created elsewhere..
	ctx := context.Background()

	eventsQueryClient := GetGoMemory(clientID).(client.EventsQueryClient)
	// TODO_IN_THIS_COMMIT: handle error...
	eventsBytesObs, err := eventsQueryClient.EventsBytes(ctx, C.GoString(query))
	if err != nil {
		*cErr = C.CString(err.Error())
		return C.int64_t(-1)
	} else {
		*cErr = C.CString("")
	}

	return C.int64_t(SetGoMemory(eventsBytesObs))
}
