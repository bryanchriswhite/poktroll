package main

import "C"
import (
	"context"

	"cosmossdk.io/depinject"

	"github.com/pokt-network/poktroll/pkg/client/block"
)

// #include <client.h>
import "C"

//export NewBlockClient
func NewBlockClient(depsRef C.GoRef, cErr **C.char) C.GoRef {
	// TODO_CONSIDERATION: Could support a version of methods which receive a go context, created elsewhere..
	ctx := context.Background()

	deps, err := GetGoMem[depinject.Config](depsRef)
	if err != nil {
		*cErr = C.CString(err.Error())
		return CNilRef
	}

	blockClient, err := block.NewBlockClient(ctx, deps)
	if err != nil {
		*cErr = C.CString(err.Error())
		return CNilRef
	}

	return SetGoMem(blockClient)
}
