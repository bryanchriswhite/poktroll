package main

// #cgo CFLAGS: -I./include
// #include <client.h>
// #include <stdint.h>
// #include <errno.h>
import "C"
import (
	"context"

	sdkclient "github.com/cosmos/cosmos-sdk/client"

	"github.com/pokt-network/poktroll/pkg/client"
)

//export NewBlockQueryClient
func NewBlockQueryClient(cometWebsocketURL *C.char, cErr **C.char) C.GoRef {
	// TODO_TECHDEBT: support opts args.
	blockQueryClient, err := sdkclient.NewClientFromNode(C.GoString(cometWebsocketURL))
	if err != nil {
		*cErr = C.CString(err.Error())
		return CNilRef
	}

	return SetGoMem(blockQueryClient)
}

//export BlockQueryClientBlock
func BlockQueryClientBlock(clientRef C.GoRef, cHeight *C.int64_t, cErr **C.char) C.GoRef {
	var height *int64
	if cHeight != nil {
		*height = int64(*cHeight)
	} else {
	}

	// TODO_CONSIDERATION: Could support a version of methods which receive a go context, created elsewhere..
	ctx := context.Background()

	blockQueryClient, err := GetGoMem[client.BlockQueryClient](clientRef)
	if err != nil {
		*cErr = C.CString(err.Error())
		return CNilRef
	}

	resultBlock, err := blockQueryClient.Block(ctx, height)
	if err != nil {
		*cErr = C.CString(err.Error())
		return CNilRef
	}

	// TODO_IN_THIS_COMMIT: return C-native struct.
	return SetGoMem(resultBlock)
}
